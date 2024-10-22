package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameCourse = "course"

type Course struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name        string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`

	Chapters []Chapter ``
}

func (*Course) TableName() string {
	return TableNameCourse
}
