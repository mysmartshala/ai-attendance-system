package services

import (
	"encoding/json"
	"ai-attendance-system/models"
	"github.com/jinzhu/gorm"
)

type StudentService struct {
	db          *gorm.DB
	faceService *FaceService
}

func NewStudentService(db *gorm.DB, faceService *FaceService) *StudentService {
	return &StudentService{
		db:          db,
		faceService: faceService,
	}
}

func (ss *StudentService) CreateStudentWithPhoto(student *models.Student, photoPath string) error {
	embedding, err := ss.faceService.GenerateEmbedding(photoPath)
	if err != nil {
		return err
	}

	embJSON, _ := json.Marshal(embedding)
	student.Embedding = string(embJSON)
	student.PhotoPath = photoPath

	return ss.db.Create(student).Error
}

func (ss *StudentService) UpdateStudentPhoto(studentID uint, photoPath string) error {
	embedding, err := ss.faceService.GenerateEmbedding(photoPath)
	if err != nil {
		return err
	}

	embJSON, _ := json.Marshal(embedding)

	return ss.db.Model(&models.Student{}).
		Where("id = ?", studentID).
		Updates(map[string]interface{}{
			"photo_path": photoPath,
			"embedding":  string(embJSON),
		}).Error
}

func (ss *StudentService) GetStudent(id uint) (*models.Student, error) {
	var student models.Student
	err := ss.db.First(&student, id).Error
	return &student, err
}

func (ss *StudentService) ListStudents(filters map[string]interface{}) ([]models.Student, error) {
	var students []models.Student
	query := ss.db.Where("is_active = ?", true)

	if course, ok := filters["course"]; ok {
		query = query.Where("course = ?", course)
	}
	if semester, ok := filters["semester"]; ok {
		query = query.Where("semester = ?", semester)
	}

	err := query.Find(&students).Error
	return students, err
}

func (ss *StudentService) DeleteStudent(id uint) error {
	return ss.db.Model(&models.Student{}).Where("id = ?", id).Update("is_active", false).Error
}

func (ss *StudentService) UpdateStudent(student *models.Student) error {
	return ss.db.Save(student).Error
}
