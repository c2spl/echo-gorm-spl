package model

import (
	"time"
)

type User struct {
	ID int

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	Name     string
	Email    string
	Password string `gorm:"->:false;<-"` //`gorm:"<-;->:false"`
}
