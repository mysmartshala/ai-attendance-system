package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"ai-attendance-system/models"
	"ai-attendance-system/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type StudentHandler struct {
	studentService *services.StudentService
	uploadPath     string
}

func NewStudentHandler(ss *services.StudentService, uploadPath string) *StudentHandler {
	return &StudentHandler{
		studentService: ss,
		uploadPath:     uploadPath,
	}
}

func (sh *StudentHandler) CreateStudent(c *gin.Context) {
	var req struct {
		RollNo   string `form:"roll_no" binding:"required"`
		Name     string `form:"name" binding:"required"`
		Course   string `form:"course" binding:"required"`
		Semester int    `form:"semester" binding:"required"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "photo is required"})
		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	filepath := filepath.Join(sh.uploadPath, "student_photos", filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save photo"})
		return
	}

	student := &models.Student{
		RollNo:   req.RollNo,
		Name:     req.Name,
		Course:   req.Course,
		Semester: req.Semester,
	}

	if err := sh.studentService.CreateStudentWithPhoto(student, filepath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func (sh *StudentHandler) ListStudents(c *gin.Context) {
	filters := make(map[string]interface{})
	if course := c.Query("course"); course != "" {
		filters["course"] = course
	}
	if semester := c.Query("semester"); semester != "" {
		if sem, err := strconv.Atoi(semester); err == nil {
			filters["semester"] = sem
		}
	}

	students, err := sh.studentService.ListStudents(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func (sh *StudentHandler) GetStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	student, err := sh.studentService.GetStudent(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (sh *StudentHandler) UpdateStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	var req struct {
		Name     string `form:"name"`
		Course   string `form:"course"`
		Semester int    `form:"semester"`
	}

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student, _ := sh.studentService.GetStudent(uint(id))
	if req.Name != "" {
		student.Name = req.Name
	}
	if req.Course != "" {
		student.Course = req.Course
	}
	if req.Semester != 0 {
		student.Semester = req.Semester
	}

	if file, err := c.FormFile("photo"); err == nil {
		filename := uuid.New().String() + filepath.Ext(file.Filename)
		filepath := filepath.Join(sh.uploadPath, "student_photos", filename)
		if err := c.SaveUploadedFile(file, filepath); err == nil {
			sh.studentService.UpdateStudentPhoto(uint(id), filepath)
		}
	}

	sh.studentService.UpdateStudent(student)
	c.JSON(http.StatusOK, student)
}

func (sh *StudentHandler) DeleteStudent(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	if err := sh.studentService.DeleteStudent(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student deleted"})
}
