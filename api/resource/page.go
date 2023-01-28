package resource

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_cms/api/models"
	"github.com/leeliwei930/notion_sdk/actions"
	"github.com/leeliwei930/notion_sdk/database/filter"
)

func GetDefaultPageResource() (*models.PageConfiguration, error) {
	pageConfigDatabaseId, uuidErr := uuid.Parse(PageConfigurationDatabaseId)
	if uuidErr != nil {
		return nil, uuidErr
	}

	cursor, queryErr := actions.QueryDatabase(pageConfigDatabaseId,
		actions.FilterWith(&filter.QueryProps{
			Property: "Name",
			RichText: &filter.Text{
				Equals: "Default",
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

	// siteConfigId, siteConfigIdParseErr := uuid.Parse(properties["Site Config"].ID)
	// if siteConfigIdParseErr != nil {
	// 	return nil, siteConfigIdParseErr
	// }
	var wg sync.WaitGroup

	landingPageRetrievedErrChan := make(chan error)
	landingPageChan := make(chan *models.LandingPage)
	wg.Add(1)

	go func() {
		defer wg.Done()

		landingPage, landingPageRetrievedErr := actions.RetrievePage(landingPageConfigId)
		if landingPageRetrievedErr != nil {
			landingPageRetrievedErrChan <- landingPageRetrievedErr
			return
		}

		var coverImageUrl string
		if len(landingPage.Properties["Cover Image"].Files) > 0 {
			coverImageUrl = landingPage.Properties["Cover Image"].Files[0].File.Url
		}

		landingPageChan <- &models.LandingPage{
			Title:               landingPage.Properties["Title"].Title[0].PlainText,
			Description:         landingPage.Properties["Description"].RichText[0].PlainText,
			CoverImage:          coverImageUrl,
			PrimaryButtonText:   landingPage.Properties["Primary Button Text"].RichText[0].PlainText,
			SecondaryButtonText: landingPage.Properties["Secondary Button Text"].RichText[0].PlainText,
			SecondaryButtonLink: landingPage.Properties["Secondary Button Link"].Url,
		}
	}()

	go func() {
		wg.Wait()
		close(landingPageChan)
		close(landingPageRetrievedErrChan)
	}()
	pageConfig := &models.PageConfiguration{}

	select {
	case landingPage := <-landingPageChan:
		pageConfig.LandingPage = landingPage
		return pageConfig, nil
	case receiveErr := <-landingPageRetrievedErrChan:
		return nil, receiveErr
	}

}
