package models

type PageConfiguration struct {
	LandingPage *LandingPage `json:"landing_page,omitempty"`
	Website     *Website     `json:"website,omitempty"`
}

type LandingPage struct {
	Title               string `json:"title,omitempty"`
	Description         string `json:"description,omitempty"`
	CoverImage          string `json:"cover_image,omitempty"`
	PrimaryButtonText   string `json:"primary_button_text,omitempty"`
	SecondaryButtonText string `json:"secondary_button_text,omitempty"`
	SecondaryButtonLink string `json:"secondary_button_link,omitempty"`
}

type Website struct {
	Name        string `json:"name,omitempty"`
	Separator   string `json:"separator,omitempty"`
	TitleFormat string `json:"title_format,omitempty"`
}
