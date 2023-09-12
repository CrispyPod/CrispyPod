package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type EpisodeStatusType int

const (
	EpisodeStatus_Draft     EpisodeStatusType = 0
	EpisodeStatus_Published EpisodeStatusType = 1
)

type Episode struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid(),primary_key"`
	Title         string
	CreateTime    int64 `gorm:"autoCreateTime"`
	PublishTime   time.Time
	Description   string
	EpisodeStatus EpisodeStatusType

	ThumbnailFileName   sql.NullString
	ThumbnailUploadName sql.NullString

	AudioFileName       sql.NullString
	AudioFileUploadName sql.NullString
	AudioFileDuration   sql.NullInt64

	UserID uuid.UUID
	User   User
}
