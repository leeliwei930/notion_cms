package resource

import (
	"github.com/google/uuid"
	"github.com/leeliwei930/notion_cms/generated/portfolio_graphql/models"
	"github.com/leeliwei930/notion_sdk/actions"
	notionmodels "github.com/leeliwei930/notion_sdk/models"
)

func GetMilestonesExperiencesResource() ([]*models.MilestonePayload, error) {
	milestoneDatabaseId, uuidErr := uuid.Parse(MilestoneDatabaseId)
	if uuidErr != nil {
		return nil, uuidErr
	}

	cursor, queryErr := actions.QueryDatabase(milestoneDatabaseId)
	if queryErr != nil {
		return nil, queryErr
	}
	milestones := []*models.MilestonePayload{}
	// fetch experience page record
	for _, milestoneRec := range cursor.Results {
		yearProp := milestoneRec.Properties["Year"].Title[0].PlainText

		themeChan, themeChanErr := FetchMilestoneTheme(milestoneRec.Properties["Theme"].Relation[0].ID)
		theme, themeErr := <-themeChan, <-themeChanErr
		if themeErr != nil {
			return nil, themeErr
		}

		experiencesChan, experiencesChanErr := FetchExperiences(milestoneRec.Properties["Experiences"].Relation)
		experiences, experiencesErr := <-experiencesChan, <-experiencesChanErr
		if experiencesErr != nil {
			return nil, experiencesErr
		}

		milestones = append(milestones, &models.MilestonePayload{
			Year:        yearProp,
			Role:        milestoneRec.Properties["Role"].RichText[0].PlainText,
			Summary:     milestoneRec.Properties["Summary"].RichText[0].PlainText,
			Theme:       &theme,
			Experiences: experiences,
		})

	}

	return milestones, nil
}

func FetchMilestoneTheme(milestoneThemeId string) (<-chan models.Theme, <-chan error) {
	themeChan := make(chan models.Theme)
	errorChan := make(chan error)

	go func() {
		defer close(themeChan)
		defer close(errorChan)
		themePage, themePageErr := actions.RetrievePage(uuid.MustParse(milestoneThemeId))
		if themePageErr != nil {
			errorChan <- themePageErr
			return
		}
		themeChan <- models.Theme{
			BackgroundColor: themePage.Properties["Background Color"].Select.Name,
			AccentColor:     themePage.Properties["Accent Color"].Select.Name,
			HeadlineColor:   themePage.Properties["Headline Color"].Select.Name,
		}
	}()

	return themeChan, errorChan
}

func FetchExperiences(experienceIds []notionmodels.RelationPropertyValue) (<-chan []*models.Experience, <-chan error) {
	experiencesChan := make(chan []*models.Experience)
	errorChan := make(chan error)

	go func() {
		defer close(experiencesChan)
		defer close(errorChan)

		var experiences []*models.Experience

		for _, experienceId := range experienceIds {
			experiencePage, experiencePageErr := actions.RetrievePage(uuid.MustParse(experienceId.ID))
			if experiencePageErr != nil {
				errorChan <- experiencePageErr
				return
			}
			experiences = append(experiences, &models.Experience{
				RawContent: experiencePage.Properties["Content"].Title[0].PlainText,
			})
		}

		experiencesChan <- experiences
	}()

	return experiencesChan, errorChan

}
