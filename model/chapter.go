package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameChapter = "chapter"

type Chapter struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name  string `gorm:"type:character varying(80);not null"`
	Order int    `gorm:"type:integer;not null"`

	CourseID uuid.UUID `gorm:"type:uuid;not null"`
	Course   Course    `gorm:"constraint:OnDelete:CASCADE;"`
	Sections []Section ``
}

func (*Chapter) TableName() string {
	return TableNameChapter
}
