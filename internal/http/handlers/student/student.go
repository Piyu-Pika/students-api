package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Piyu-Pika/students-api/internal/storage"
	"github.com/Piyu-Pika/students-api/internal/types"
	"github.com/Piyu-Pika/students-api/internal/utils/responce"
	"github.com/go-playground/validator/v10"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("create student")

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			responce.WriteJson(w, http.StatusBadRequest, responce.GenralError(err))
			return

		}
		if err != nil {
			responce.WriteJson(w, http.StatusBadRequest, responce.GenralError(err))
			return
		}

		//request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			responce.WriteJson(w, http.StatusBadRequest, responce.ValidationError(validateErrs))
			return
		}

		lastid, err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Phone,
			student.Age,
		)

		slog.Info("Student created", slog.String("id", fmt.Sprintf("%d", lastid)))
		if err != nil {
			responce.WriteJson(w, http.StatusInternalServerError, responce.GenralError(err))
			return
		}

		w.Write([]byte("Welcome to Students API \n"))
		fmt.Fprintf(w, "Hello World")

		responce.WriteJson(w, http.StatusCreated, map[string]string{"id": fmt.Sprintf("%d", lastid)})
	}
}

func Getbyid(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		slog.Info("Get by id",slog.String("id",id))
		intid, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			slog.Error("Error parsing id", slog.String("id", id), slog.String("error", err.Error()))
			responce.WriteJson(w, http.StatusInternalServerError, responce.GenralError(err))
			return
		}

		student, err := storage.GetStudentByID(intid)
		if err != nil {
			slog.Error("Error getting student by id", slog.String("id", id), slog.String("error", err.Error()))
			responce.WriteJson(w, http.StatusInternalServerError, responce.GenralError(err))
			return
		}
		responce.WriteJson(w, http.StatusOK, student)  
		 


		

}
}
