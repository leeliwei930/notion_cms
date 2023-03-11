package resource

import (
	"errors"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql/models"
	"github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/database/filter"
)

func GetDefaultPageResource() (*models.PageConfigurationPayload, error) {
	pageConfigDatabaseId, uuidErr := uuid.Parse(PageConfigurationDatabaseId)
	if uuidErr != nil {
		return nil, uuidErr
	}

	defaultConfig := "Default"
	cursor, queryErr := actions.QueryDatabase(pageConfigDatabaseId,
		actions.FilterWith(&filter.QueryProps{
			Property: "Name",
			RichText: &filter.Text{
				Equals: defaultConfig,
			},
		}))
	if queryErr != nil {
		return nil, queryErr
	}

	if len(cursor.Results) == 0 {
		return nil, errors.New("no default configuration set, please ensure there is a record \"Default\" value in the name column")
	}

	properties := cursor.Results[0].Properties
	landingPageConfigId, landingPageConfigIdParseErr := uuid.Parse(properties["Active Landing Page"].Relation[0].ID)
	if landingPageConfigIdParseErr != nil {
		return nil, landingPageConfigIdParseErr
	}

	siteConfigId, siteConfigIdParseErr := uuid.Parse(properties["Site Config"].Relation[0].ID)
	if siteConfigIdParseErr != nil {
		return nil, siteConfigIdParseErr
	}

	websiteChan, websiteErrChan := FetchWebsiteConfig(siteConfigId)
	landingPageChan, landingPageErrChan := FetchLandingPageConfig(landingPageConfigId)

	website, websiteErr := <-websiteChan, <-websiteErrChan
	if websiteErr != nil {
		return nil, websiteErr
	}
	landingPage, landingPageErr := <-landingPageChan, <-landingPageErrChan
	if landingPageErr != nil {
		return nil, landingPageErr
	}

	return &models.PageConfigurationPayload{
		LandingPage: &landingPage,
		Website:     &website,
	}, nil

}

func FetchWebsiteConfig(siteConfigId uuid.UUID) (<-chan models.Website, <-chan error) {
	errorChan := make(chan error)
	websiteChan := make(chan models.Website)
	go func() {
		defer close(errorChan)
		defer close(websiteChan)
		siteConfigPage, siteConfigRetrievedErr := actions.RetrievePage(siteConfigId)
		if siteConfigRetrievedErr != nil {
			errorChan <- siteConfigRetrievedErr
			return
		}

		websiteChan <- models.Website{
			Name:        siteConfigPage.Properties["Name"].Title[0].PlainText,
			Separator:   siteConfigPage.Properties["Separator"].RichText[0].PlainText,
			TitleFormat: siteConfigPage.Properties["Title Format"].Select.Name,
		}
	}()
	return websiteChan, errorChan

}

func FetchLandingPageConfig(landingPageConfigId uuid.UUID) (<-chan models.LandingPage, <-chan error) {
	errorChan := make(chan error)
	landingPageChan := make(chan models.LandingPage)
	go func() {
		defer close(errorChan)
		defer close(landingPageChan)

		landingPage, landingPageRetrievedErr := actions.RetrievePage(landingPageConfigId)
		if landingPageRetrievedErr != nil {
			errorChan <- landingPageRetrievedErr
			return
		}

		var coverImageUrl string

		if len(landingPage.Properties["Cover Image"].Files) > 0 {
			coverImageUrl = landingPage.Properties["Cover Image"].Files[0].File.Url
		}

		landingPageChan <- models.LandingPage{
			Title:               landingPage.Properties["Title"].Title[0].PlainText,
			Description:         landingPage.Properties["Description"].RichText[0].PlainText,
			CoverImage:          coverImageUrl,
			PrimaryButtonText:   landingPage.Properties["Primary Button Text"].RichText[0].PlainText,
			SecondaryButtonText: landingPage.Properties["Secondary Button Text"].RichText[0].PlainText,
			SecondaryButtonLink: landingPage.Properties["Secondary Button Link"].Url,
		}
	}()
	return landingPageChan, errorChan
}
