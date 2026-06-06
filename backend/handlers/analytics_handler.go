package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AnalyticsHandler struct {
	db *gorm.DB
}

func NewAnalyticsHandler(db *gorm.DB) *AnalyticsHandler {
	return &AnalyticsHandler{db: db}
}

func (ah *AnalyticsHandler) GetDashboard(c *gin.Context) {
	var totalStudents int
	var todayAttendance int
	var todayAbsence int

	today := time.Now().Format("2006-01-02")

	ah.db.Table("students").Where("is_active = ?", true).Count(&totalStudents)
	ah.db.Table("attendance").Where("DATE(attendance_date) = ? AND status = ?", today, "Present").Count(&todayAttendance)
	ah.db.Table("attendance").Where("DATE(attendance_date) = ? AND status = ?", today, "Absent").Count(&todayAbsence)

	var attendancePercentage float32
	if todayAttendance+todayAbsence > 0 {
		attendancePercentage = float32(todayAttendance) / float32(todayAttendance+todayAbsence) * 100
	}

	c.JSON(http.StatusOK, gin.H{
		"total_students":        totalStudents,
		"todays_attendance":     todayAttendance,
		"todays_absence":        todayAbsence,
		"attendance_percentage": attendancePercentage,
	})
}

func (ah *AnalyticsHandler) GetCourseWiseAttendance(c *gin.Context) {
	type CourseStats struct {
		Course     string
		Percentage float32
	}

	var results []CourseStats

	query := `
	SELECT 
		s.course,
		ROUND(SUM(CASE WHEN a.status = 'Present' THEN 1 ELSE 0 END) / 
		CAST(COUNT(*) AS FLOAT) * 100, 2) AS percentage
	FROM students s
	LEFT JOIN attendance a ON s.id = a.student_id
	WHERE s.is_active = 1
	GROUP BY s.course
	`

	ah.db.Raw(query).Scan(&results)
	c.JSON(http.StatusOK, results)
}

func (ah *AnalyticsHandler) GetStudentWiseAnalytics(c *gin.Context) {
	studentID := c.Param("student_id")

	type StudentStats struct {
		StudentID  uint
		RollNo     string
		Name       string
		Present    int
		Absent     int
		Percentage float32
	}

	var stats StudentStats

	query := `
	SELECT 
		s.id AS student_id,
		s.roll_no,
		s.name,
		SUM(CASE WHEN a.status = 'Present' THEN 1 ELSE 0 END) AS present,
		SUM(CASE WHEN a.status = 'Absent' THEN 1 ELSE 0 END) AS absent,
		ROUND(SUM(CASE WHEN a.status = 'Present' THEN 1 ELSE 0 END) / 
		CAST(COUNT(*) AS FLOAT) * 100, 2) AS percentage
	FROM students s
	LEFT JOIN attendance a ON s.id = a.student_id
	WHERE s.id = ?
	GROUP BY s.id, s.roll_no, s.name
	`

	ah.db.Raw(query, studentID).Scan(&stats)
	c.JSON(http.StatusOK, stats)
}
