package main

import (
	"log"
	"os"

	"ai-attendance-system/config"
	"ai-attendance-system/database"
	"ai-attendance-system/handlers"
	"ai-attendance-system/middleware"
	"ai-attendance-system/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.Close()

	faceService := services.NewFaceService(cfg.FaceServiceURL)
	studentService := services.NewStudentService(db, faceService)
	attendanceService := services.NewAttendanceService(db, faceService, cfg.SimilarityThreshold)

	studentHandler := handlers.NewStudentHandler(studentService, cfg.UploadPath)
	attendanceHandler := handlers.NewAttendanceHandler(attendanceService, cfg.UploadPath)
	teacherHandler := handlers.NewTeacherHandler(db)
	analyticsHandler := handlers.NewAnalyticsHandler(db)

	os.MkdirAll(cfg.UploadPath+"/student_photos", 0755)
	os.MkdirAll(cfg.UploadPath+"/classroom_photos", 0755)

	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/login", teacherHandler.Login)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/students", studentHandler.CreateStudent)
		protected.GET("/students", studentHandler.ListStudents)
		protected.GET("/students/:id", studentHandler.GetStudent)
		protected.PUT("/students/:id", studentHandler.UpdateStudent)
		protected.DELETE("/students/:id", studentHandler.DeleteStudent)

		protected.POST("/attendance/process", attendanceHandler.ProcessAttendance)
		protected.GET("/attendance/report", attendanceHandler.GetAttendanceReport)

		protected.GET("/analytics/dashboard", analyticsHandler.GetDashboard)
		protected.GET("/analytics/course-wise", analyticsHandler.GetCourseWiseAttendance)
		protected.GET("/analytics/student-wise/:student_id", analyticsHandler.GetStudentWiseAnalytics)
	}

	port := ":" + cfg.ServerPort
	log.Printf("Starting server on %s", port)
	router.Run(port)
}
