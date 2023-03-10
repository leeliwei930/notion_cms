// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"time"
)

type EducationPathwayPayload struct {
	Title         string     `json:"title"`
	InstituteName string     `json:"instituteName"`
	StudyArea     string     `json:"studyArea"`
	Icon          string     `json:"icon"`
	Image         string     `json:"image"`
	Location      string     `json:"location"`
	CommencedOn   *time.Time `json:"commencedOn"`
	CompletedOn   *time.Time `json:"completedOn"`
}

type EducationQuery struct {
	All []*EducationPathwayPayload `json:"all"`
}

type Experience struct {
	RawContent string `json:"rawContent"`
}

type LandingPage struct {
	Title               string `json:"title"`
	Description         string `json:"description"`
	CoverImage          string `json:"coverImage"`
	PrimaryButtonText   string `json:"primaryButtonText"`
	SecondaryButtonText string `json:"secondaryButtonText"`
	SecondaryButtonLink string `json:"secondaryButtonLink"`
}

type MilestonePayload struct {
	Year        string        `json:"year"`
	Summary     string        `json:"summary"`
	Role        string        `json:"Role"`
	Theme       *Theme        `json:"theme"`
	Experiences []*Experience `json:"experiences"`
}

type MilestoneQuery struct {
	All []*MilestonePayload `json:"all"`
}

type PageConfigurationPayload struct {
	LandingPage *LandingPage `json:"landingPage"`
	Website     *Website     `json:"website"`
}

type Theme struct {
	BackgroundColor string `json:"backgroundColor"`
	AccentColor     string `json:"accentColor"`
	HeadlineColor   string `json:"headlineColor"`
}

type Website struct {
	Name        string `json:"name"`
	Separator   string `json:"separator"`
	TitleFormat string `json:"titleFormat"`
}

type WebsiteQuery struct {
	Config *PageConfigurationPayload `json:"config"`
}
