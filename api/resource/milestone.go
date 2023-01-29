package resource

import (
	"github.com/google/uuid"
	"github.com/leeliwei930/notion_cms/api/models"
	"github.com/leeliwei930/notion_sdk/actions"
	notionModels "github.com/leeliwei930/notion_sdk/models"
)

func GetMilestonesExperiencesResource() ([]*models.Milestone, error) {
	milestoneDatabaseId, uuidErr := uuid.Parse(MilestoneDatabaseId)
	if uuidErr != nil {
		return nil, uuidErr
	}

	cursor, queryErr := actions.QueryDatabase(milestoneDatabaseId)
	if queryErr != nil {
		return nil, queryErr
	}
	milestonesRes := []*models.Milestone{}
	milestones := []*models.Milestone{}
	milestoneThemes := make(map[string]string)
	milestoneExperiences := make(map[string][]notionModels.RelationPropertyValue)
	// fetch experience page record
	for _, milestoneRec := range cursor.Results {
		yearProp := milestoneRec.Properties["Year"].Title[0].PlainText
		milestones = append(milestones, &models.Milestone{
			Year:    yearProp,
			Role:    milestoneRec.Properties["Role"].RichText[0].PlainText,
			Summary: milestoneRec.Properties["Summary"].RichText[0].PlainText,
		})
		milestoneThemes[yearProp] = milestoneRec.Properties["Theme"].Relation[0].ID
		milestoneExperiences[yearProp] = milestoneRec.Properties["Experiences"].Relation
	}

	// fetch theme record
	for year, themeId := range milestoneThemes {
		themePage, themePageErr := actions.RetrievePage(uuid.MustParse(themeId))
		if themePageErr != nil {
			return milestonesRes, themePageErr
		}
		milestone := findMilestoneByYear(year, milestones)
		milestone.Theme = &models.Theme{
			BackgroundColor: themePage.Properties["Background Color"].Select.Name,
			AccentColor:     themePage.Properties["Accent Color"].Select.Name,
			HeadlineColor:   themePage.Properties["Headline Color"].Select.Name,
		}
	}

	for year, experienceIds := range milestoneExperiences {
		experiences := []*models.Experience{}
		for _, experienceId := range experienceIds {
			experiencePage, experiencePageErr := actions.RetrievePage(uuid.MustParse(experienceId.ID))
			if experiencePageErr != nil {
				return milestonesRes, experiencePageErr
			}

			experiences = append(experiences, &models.Experience{
				RawContent: experiencePage.Properties["Content"].Title[0].PlainText,
			})

		}
		milestone := findMilestoneByYear(year, milestones)
		milestone.Experience = append(milestone.Experience, experiences...)
	}

	milestonesRes = append(milestonesRes, milestones...)

	return milestonesRes, nil
}

func findMilestoneByYear(year string, milestones []*models.Milestone) *models.Milestone {
	for _, milestone := range milestones {
		if milestone.Year == year {
			return milestone
		}
	}
	return nil
}
