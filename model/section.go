package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TableNameSection = "section"

type Section struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name    string `gorm:"type:character varying(80);not null"`
	Order   int    `gorm:"type:integer;not null"`
	Title   string `gorm:"type:text;not null"`
	Content string `gorm:"type:text;not null"`

	ChapterID uuid.UUID `gorm:"type:uuid;not null"`
	Chapter   Chapter   `gorm:"constraint:OnDelete:CASCADE;"`
}

func (*Section) TableName() string {
	return TableNameSection
}
