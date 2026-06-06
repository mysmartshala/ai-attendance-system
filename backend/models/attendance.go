package models

import (
	"time"
)

type Attendance struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	StudentID      uint      `json:"student_id"`
	Student        Student   `gorm:"foreignKey:StudentID" json:"student"`
	AttendanceDate time.Time `gorm:"type:date" json:"attendance_date"`
	Status         string    `json:"status"`
	Confidence     float32   `json:"confidence"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Attendance) TableName() string {
	return "attendance"
}
