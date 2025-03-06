package responce

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func GenralError(err error) Response {
	return Response{
		Status:  "error",
		Message: err.Error(),
	}

}

func ValidationError(errs validator.ValidationErrors) Response {
	var errmessage []string
	for _, e := range errs {
		switch e.Tag() {
		case "required":
			errmessage = append(errmessage, fmt.Sprintf("%s is required field", e.Field()))
		default:
			errmessage = append(errmessage, fmt.Sprintf("%s %s", e.Field(), e.ActualTag))
		}
	}
	return Response{
		Status:  "error",
		Message: strings.Join(errmessage, ","),
	}
}
