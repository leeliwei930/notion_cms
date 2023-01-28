package resource

import (
	"errors"

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

	landingPageConfigId, landingPageConfigIdParseErr := uuid.Parse(properties["Active Landing Page"].ID)
	if landingPageConfigIdParseErr != nil {
		return nil, landingPageConfigIdParseErr
	}

	siteConfigId, siteConfigIdParseErr := uuid.Parse(properties["Site Config"].ID)
	if siteConfigIdParseErr != nil {
		return nil, siteConfigIdParseErr
	}

	// TODO(leeliwei930): "Query landingPageConfig and siteConfigPage properties using GetPage action"

	pageConfig := &models.PageConfiguration{}

	return pageConfig, nil
}
