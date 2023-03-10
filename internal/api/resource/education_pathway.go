package resource

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql/models"
	"github.com/leeliwei930/notion_sdk/actions"
)

func GetEducationPathwayResource() ([]*models.EducationPathwayPayload, error) {
	educationPathwayDatabaseId, uuidErr := uuid.Parse(EducationPathwayDatabaseId)
	if uuidErr != nil {
		return nil, uuidErr
	}

	cursor, queryErr := actions.QueryDatabase(educationPathwayDatabaseId)
	if queryErr != nil {
		return nil, queryErr
	}

	if len(cursor.Results) == 0 {
		return nil, fmt.Errorf("no database block contains in the %s", educationPathwayDatabaseId.String())
	}

	educationPathway := []*models.EducationPathwayPayload{}
	for _, result := range cursor.Results {
		properties := result.Properties
		educationPathway = append(educationPathway, &models.EducationPathwayPayload{
			Title:         properties["Title"].Title[0].PlainText,
			InstituteName: properties["Institute Name"].RichText[0].PlainText,
			StudyArea:     properties["Study"].RichText[0].PlainText,
			Icon:          properties["Icon"].Files[0].File.Url,
			Image:         properties["Image"].Files[0].File.Url,
			Location:      properties["Location"].RichText[0].PlainText,
			CommencedOn:   &properties["Commenced On"].Date.Start.Time,
			CompletedOn:   &properties["Completed On"].Date.Start.Time,
		})
	}

	return educationPathway, nil
}
