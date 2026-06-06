package models

import (
	"time"
)

type Student struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	RollNo    string    `gorm:"uniqueIndex" json:"roll_no"`
	Name      string    `json:"name"`
	Course    string    `json:"course"`
	Semester  int       `json:"semester"`
	PhotoPath string    `json:"photo_path"`
	Embedding string    `gorm:"type:longtext" json:"embedding"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Student) TableName() string {
	return "students"
}
