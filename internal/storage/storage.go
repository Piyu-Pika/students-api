package storage

type Storage interface {
	CreateStudent(name string, email string, phone string, age int) (int64, error)

}