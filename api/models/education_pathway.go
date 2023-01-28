package models

import "time"

type EducationPathway struct {
	Title         string     `json:"title,omitempty"`
	InstituteName string     `json:"institute,omitempty"`
	StudyArea     string     `json:"study,omitempty"`
	Icon          string     `json:"icon_url,omitempty"`
	Image         string     `json:"image,omitempty"`
	Location      string     `json:"location,omitempty"`
	CommencedOn   *time.Time `json:"commenced_on,omitempty"`
	CompletedOn   *time.Time `json:"completed_on,omitempty"`
}
