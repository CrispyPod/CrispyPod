package models

import (
	"time"

	"crispypod.com/crispypod/graph/model"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	CreateTime  time.Time
	Email       string
	UserName    string
	Password    string
	DisplayName string
	IsAdmin     bool
}

func (u *User) ToGQLUser() *model.User {
	return &model.User{
		ID:          u.ID.String(),
		Email:       u.Email,
		CreateTime:  int(u.CreateTime.Unix()),
		UserName:    u.UserName,
		DisplayName: u.DisplayName,
		IsAdmin:     u.IsAdmin,
	}
}
