package storage

import "github.com/Piyu-Pika/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, phone string, age int) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
