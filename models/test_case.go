package models

import (
	"time"

	"gorm.io/gorm"
)

type TestCase struct {
	ID              uint           `gorm:"primarykey;" json:"id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	Identity        string         `gorm:"column:identity;type:varchar(255);" json:"identity"`
	ProblemIdentity string         `gorm:"column:problem_identity;type:varchar(255);" json:"problem_identity"`
	Input           string         `gorm:"column:input;type:varchar(255);" json:"input"`
	Output          string         `gorm:"column:output;type:varchar(255);" json:"output"`
}

func (table *TestCase) TableName() string {
	return "test_case"
}
