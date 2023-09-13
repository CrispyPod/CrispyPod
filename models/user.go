package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	CreateTime  int64     `gorm:"autoCreateTime"`
	Email       string
	UserName    string
	Password    string
	DisplayName string
	IsAdmin     bool
}
