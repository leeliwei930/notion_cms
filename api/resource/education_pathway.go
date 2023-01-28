package resource

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/leeliwei930/notion_cms/api/models"
	"github.com/leeliwei930/notion_sdk/actions"
)

func GetEducationPathwayResource() ([]models.EducationPathway, error) {
	educationPathwayDatabaseId, uuidErr := uuid.Parse("9ff25e1831bc475a9ff89b833fb536c3")
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

	educationPathway := []models.EducationPathway{}
	for _, result := range cursor.Results {
		properties := result.Properties
		educationPathway = append(educationPathway, models.EducationPathway{
			Title:         properties["Title"].Title[0].PlainText,
			InstituteName: properties["Institute Name"].RichText[0].PlainText,
			StudyArea:     properties["Study"].RichText[0].PlainText,
			Icon:          properties["Icon"].Files[0].File.Url,
			Image:         properties["Image"].Files[0].File.Url,
			CommencedOn:   &properties["Commenced On"].Date.Start.Time,
			CompletedOn:   &properties["Completed On"].Date.Start.Time,
		})
	}

	return educationPathway, nil
}
