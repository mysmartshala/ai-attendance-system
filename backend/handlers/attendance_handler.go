package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"ai-attendance-system/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttendanceHandler struct {
	attendanceService *services.AttendanceService
	uploadPath        string
}

func NewAttendanceHandler(as *services.AttendanceService, uploadPath string) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: as,
		uploadPath:        uploadPath,
	}
}

func (ah *AttendanceHandler) ProcessAttendance(c *gin.Context) {
	var req struct {
		Course   string `form:"course" binding:"required"`
		Semester int    `form:"semester" binding:"required"`
		Date     string `form:"date"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendanceDate := time.Now()
	if req.Date != "" {
		parsed, err := time.Parse("2006-01-02", req.Date)
		if err == nil {
			attendanceDate = parsed
		}
	}

	file, err := c.FormFile("classroom_photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "classroom_photo is required"})
		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	filepath := filepath.Join(ah.uploadPath, "classroom_photos", filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save photo"})
		return
	}

	result, err := ah.attendanceService.ProcessAttendance(filepath, req.Course, req.Semester, attendanceDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ah *AttendanceHandler) GetAttendanceReport(c *gin.Context) {
	course := c.Query("course")
	semester := c.Query("semester")
	startDate := c.DefaultQuery("start_date", time.Now().AddDate(0, 0, -30).Format("2006-01-02"))
	endDate := c.DefaultQuery("end_date", time.Now().Format("2006-01-02"))

	if course == "" || semester == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course and semester required"})
		return
	}

	sem, _ := strconv.Atoi(semester)
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	report, err := ah.attendanceService.GetAttendanceReport(course, sem, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, report)
}
