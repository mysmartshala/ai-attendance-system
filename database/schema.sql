-- Create Database
CREATE DATABASE IF NOT EXISTS attendance;
USE attendance;

-- Students Table
CREATE TABLE students (
  id INT AUTO_INCREMENT PRIMARY KEY,
  roll_no VARCHAR(50) UNIQUE NOT NULL,
  name VARCHAR(100) NOT NULL,
  course VARCHAR(50) NOT NULL,
  semester INT NOT NULL,
  photo_path VARCHAR(255),
  embedding LONGTEXT,
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_course_semester (course, semester),
  INDEX idx_is_active (is_active),
  INDEX idx_roll_no (roll_no)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Attendance Table
CREATE TABLE attendance (
  id INT AUTO_INCREMENT PRIMARY KEY,
  student_id INT NOT NULL,
  attendance_date DATE NOT NULL,
  status VARCHAR(20) NOT NULL DEFAULT 'Absent',
  confidence FLOAT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE,
  INDEX idx_student_date (student_id, attendance_date),
  INDEX idx_attendance_date (attendance_date),
  INDEX idx_status (status),
  UNIQUE KEY unique_attendance (student_id, attendance_date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Teachers Table
CREATE TABLE teachers (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  INDEX idx_username (username),
  INDEX idx_is_active (is_active)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert Default Teachers
INSERT INTO teachers (username, password) VALUES ('teacher', 'teacher123');
INSERT INTO teachers (username, password) VALUES ('admin', 'admin123');

-- Sample Students for Testing
INSERT INTO students (roll_no, name, course, semester) VALUES 
('BCA001', 'Rahul Kumar', 'BCA', 3),
('BCA002', 'Anjali Patil', 'BCA', 3),
('BCA003', 'Mahesh Rao', 'BCA', 3),
('BCA004', 'Priya Singh', 'BCA', 3),
('BCA005', 'Amit Sharma', 'BCA', 3),
('BCom001', 'Neha Verma', 'BCom', 2),
('BCom002', 'Rohan Gupta', 'BCom', 2),
('BCom003', 'Sarah Khan', 'BCom', 2),
('BCom004', 'Vikram Singh', 'BCom', 2),
('BBA001', 'Deepika Patel', 'BBA', 4),
('BBA002', 'Arjun Reddy', 'BBA', 4),
('BBA003', 'Isha Joshi', 'BBA', 4),
('BBA004', 'Rajesh Kumar', 'BBA', 4);
