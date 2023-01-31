package models

type Milestone struct {
	Experiences []*Experience `json:"experiences,omitempty"`
	Year        string        `json:"year,omitempty"`
	Summary     string        `json:"summary,omitempty"`
	Role        string        `json:"role,omitempty"`
	Theme       *Theme        `json:"theme,omitempty"`
}

type Theme struct {
	BackgroundColor string `json:"background_color,omitempty"`
	AccentColor     string `json:"accent_color,omitempty"`
	HeadlineColor   string `json:"headline_color,omitempty"`
}

type Experience struct {
	RawContent string `json:"raw_content,omitempty"`
}
