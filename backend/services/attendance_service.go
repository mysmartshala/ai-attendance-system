package services

import (
	"encoding/json"
	"fmt"
	"time"
	"ai-attendance-system/models"
	"github.com/jinzhu/gorm"
)

type AttendanceService struct {
	db          *gorm.DB
	faceService *FaceService
	threshold   float32
}

type AttendanceResult struct {
	TotalStudents int                    `json:"total_students"`
	Detected      int                    `json:"detected"`
	Present       int                    `json:"present"`
	Unknown       int                    `json:"unknown"`
	Absent        int                    `json:"absent"`
	MatchResults  []MatchResultWithImage `json:"match_results"`
}

type MatchResultWithImage struct {
	FaceIndex   int        `json:"face_index"`
	StudentID   uint       `json:"student_id"`
	StudentName string     `json:"student_name"`
	RollNo      string     `json:"roll_no"`
	Confidence  float32    `json:"confidence"`
	BBox        [4]float32 `json:"bbox"`
	IsMatched   bool       `json:"is_matched"`
}

func NewAttendanceService(db *gorm.DB, faceService *FaceService, threshold float32) *AttendanceService {
	return &AttendanceService{
		db:          db,
		faceService: faceService,
		threshold:   threshold,
	}
}

func (as *AttendanceService) ProcessAttendance(classroomImagePath string, course string, semester int, attendanceDate time.Time) (*AttendanceResult, error) {
	var students []models.Student
	as.db.Where("course = ? AND semester = ? AND is_active = ?", course, semester, true).Find(&students)

	result := &AttendanceResult{
		TotalStudents: len(students),
		MatchResults:  []MatchResultWithImage{},
	}

	faces, err := as.faceService.DetectFaces(classroomImagePath)
	if err != nil {
		return result, err
	}

	result.Detected = len(faces)

	studentEmbeddings := make(map[uint][]float32)
	for _, student := range students {
		var emb []float32
		if student.Embedding != "" {
			json.Unmarshal([]byte(student.Embedding), &emb)
			studentEmbeddings[student.ID] = emb
		}
	}

	matchedStudentIDs := make(map[uint]float32)

	for faceIdx, face := range faces {
		bestMatch := MatchResultWithImage{
			FaceIndex: faceIdx,
			BBox:      face.BBox,
			IsMatched: false,
		}

		bestScore := float32(0)
		var bestStudentID uint

		for studentID, studentEmb := range studentEmbeddings {
			similarity := as.faceService.CosineSimilarity(face.Embedding, studentEmb)
			if similarity > bestScore && similarity >= as.threshold {
				bestScore = similarity
				bestStudentID = studentID
			}
		}

		if bestScore >= as.threshold {
			for _, student := range students {
				if student.ID == bestStudentID {
					bestMatch.StudentID = student.ID
					bestMatch.StudentName = student.Name
					bestMatch.RollNo = student.RollNo
					bestMatch.Confidence = bestScore
					bestMatch.IsMatched = true
					matchedStudentIDs[student.ID] = bestScore
					break
				}
			}
			result.Present++
		} else {
			result.Unknown++
		}

		result.MatchResults = append(result.MatchResults, bestMatch)
	}

	for studentID, confidence := range matchedStudentIDs {
		attendance := models.Attendance{
			StudentID:      studentID,
			AttendanceDate: attendanceDate,
			Status:         "Present",
			Confidence:     confidence,
		}
		as.db.Create(&attendance)
	}

	result.Absent = result.TotalStudents - result.Present

	return result, nil
}

func (as *AttendanceService) GetAttendanceReport(course string, semester int, startDate, endDate time.Time) ([]map[string]interface{}, error) {
	var students []models.Student
	as.db.Where("course = ? AND semester = ? AND is_active = ?", course, semester, true).Find(&students)

	var report []map[string]interface{}

	for _, student := range students {
		var attendanceRecords []models.Attendance
		as.db.Where("student_id = ? AND attendance_date BETWEEN ? AND ?", student.ID, startDate, endDate).Find(&attendanceRecords)

		present := 0
		for _, att := range attendanceRecords {
			if att.Status == "Present" {
				present++
			}
		}

		percentage := float32(0)
		if len(attendanceRecords) > 0 {
			percentage = (float32(present) / float32(len(attendanceRecords))) * 100
		}

		report = append(report, map[string]interface{}{
			"student_id": student.ID,
			"roll_no":    student.RollNo,
			"name":       student.Name,
			"present":    present,
			"absent":     len(attendanceRecords) - present,
			"total":      len(attendanceRecords),
			"percentage": fmt.Sprintf("%.2f", percentage),
		})
	}

	return report, nil
}
