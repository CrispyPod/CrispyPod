package models

import "github.com/google/uuid"

type SiteConfig struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key"`
	SiteName        string
	SiteDescription string
	SiteUrl         string
}