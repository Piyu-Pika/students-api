package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Piyu-Pika/students-api/internal/config"
	"github.com/Piyu-Pika/students-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		phone TEXT NOT NULL,
		age INTEGER NOT NULL
	)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil

}

func (s *Sqlite) CreateStudent(name string, email string, phone string, age int) (int64, error) {

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, phone, age) VALUES (?   , ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, phone, age)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastid, nil

}

func (s *Sqlite) GetStudentByID(id int64) (types.Student, error) {
	// student := types.Student{}
	stmt ,err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err =stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Age)  
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}
		return types.Student{}, err
	}
	return student, nil
} 

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt ,err := s.Db.Prepare("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student
	for rows.Next() {
		var student types.Student
		err = rows.Scan(&student.ID, &student.Name, &student.Email, &student.Phone, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
