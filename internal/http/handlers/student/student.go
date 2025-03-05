package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Piyu-Pika/students-api/internal/types"
	"github.com/Piyu-Pika/students-api/internal/utils/responce"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("create student")

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			responce.WriteJson(w, http.StatusBadRequest, responce.GenralError(err))
			return

		}

		w.Write([]byte("Welcome to Students API \n"))
		fmt.Fprintf(w, "Hello World")

		responce.WriteJson(w, http.StatusCreated, map[string]string{"message": "Student created successfully"})
	}
}
