# AI-Powered Classroom Attendance System

A prototype system for automated student attendance using face recognition with InsightFace AI.

## Project Structure

```
ai-attendance-system/
├── backend/                 # Go Gin API Server
│   ├── main.go
│   ├── go.mod
│   ├── Dockerfile
│   ├── config/
│   │   └── config.go
│   ├── models/
│   │   ├── student.go
│   │   ├── attendance.go
│   │   └── teacher.go
│   ├── handlers/
│   │   ├── student_handler.go
│   │   ├── attendance_handler.go
│   │   ├── teacher_handler.go
│   │   └── analytics_handler.go
│   ├── services/
│   │   ├── face_recognition_service.go
│   │   ├── attendance_service.go
│   │   └── student_service.go
│   ├── middleware/
│   │   └── auth.go
│   └── database/
│       └── db.go
├── frontend/                # PHP + Bootstrap UI
│   ├── config.php
│   ├── public/
│   │   ├── index.php
│   │   ├── admin/
│   │   ├── teacher/
│   │   └── auth/
│   ├── includes/
│   └── assets/
├── database/
│   └── schema.sql
├── face-recognition-service/
│   ├── server.py
│   ├── requirements.txt
│   └── Dockerfile
├── docker-compose.yml
├── .env.example
├── .gitignore
└── docs/
    └── API_DOCUMENTATION.md
```

## Quick Start

### Prerequisites
- Docker & Docker Compose
- Or: Go 1.19+, PHP 7.4+, MySQL 8.0+, Python 3.8+

### Setup with Docker

```bash
# Clone repository
git clone https://github.com/mysmartshala/ai-attendance-system.git
cd ai-attendance-system

# Copy environment file
cp .env.example .env

# Start services
docker-compose up -d

# Initialize database
docker exec ai-attendance-mysql mysql -u root -proot attendance < database/schema.sql

# Backend runs on: http://localhost:8080
# Frontend runs on: http://localhost
```

## Default Credentials

### Teacher Portal
- **URL**: http://localhost/teacher/dashboard.php
- **Username**: teacher
- **Password**: teacher123

### Admin Portal
- **URL**: http://localhost/admin/dashboard.php
- **Username**: admin
- **Password**: admin123

## API Endpoints

### Authentication
- `POST /api/auth/login` - Teacher login

### Students
- `POST /api/students` - Create student with photo
- `GET /api/students` - List students
- `GET /api/students/:id` - Get student details
- `PUT /api/students/:id` - Update student
- `DELETE /api/students/:id` - Delete student

### Attendance
- `POST /api/attendance/process` - Process classroom photo
- `GET /api/attendance/report` - Get attendance report

### Analytics
- `GET /api/analytics/dashboard` - Dashboard statistics
- `GET /api/analytics/course-wise` - Course-wise attendance
- `GET /api/analytics/student-wise/:studentId` - Student analytics

## Features

✅ Student Management (CRUD)
✅ Face Recognition & Embedding
✅ AI-Powered Attendance Detection
✅ Real-time Face Matching
✅ Confidence Scoring
✅ Analytics Dashboard
✅ Mobile Camera Support
✅ Teacher & Admin Portals

## Technology Stack

| Component | Technology |
|-----------|------------|
| Backend | Go (Gin) |
| Frontend | PHP + Bootstrap 5 |
| Database | MySQL 8.0 |
| Face Recognition | InsightFace (Python) |
| Authentication | JWT |
| Containerization | Docker |

## Development Timeline

| Task | Time | Status |
|------|------|--------|
| Project Setup | 1 Day | ✅ Done |
| Backend APIs | 1 Day | ✅ Done |
| Student Module | 1 Day | ✅ Done |
| Face Recognition | 2 Days | ✅ Done |
| Attendance Module | 1 Day | ✅ Done |
| Analytics & Frontend | 1 Day | ✅ Done |
| Testing | 1 Day | ⏳ Pending |
| **Total** | **~8 Days** | |

## License

MIT License
