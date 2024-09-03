package models

import (
	"time"
)

type WorkExperience struct {
	Name         string      `json:"name"`
	Position     string      `json:"position"`
	Location     string      `json:"location"`
	Achievements []string    `json:"achievements"`
	Date         Date        `json:"date"`
	TechStacks   []TechStack `json:"tech_stacks"`
}

type Project struct {
	Name        string      `json:"name"`
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Image       string      `json:"image_url"`
	TechStacks  []TechStack `json:"tech_stacks"`
}

type Date struct {
	Start time.Time   `json:"start"`
	End   interface{} `json:"end"`
}

type TechStack struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Link string `json:"link"`
}
