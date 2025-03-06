package types

type Student struct {
	ID      int64
	Name  string `validate:"required"`
	Email string `validate:"required"`
	Phone string `validate:"required"`
	Age   int    `validate:"required"`
}
