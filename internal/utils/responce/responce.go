package responce

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string
	Message string
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func GenralError(err error) Response{
	return Response{
		Status:  "error",
		Message: err.Error(),
	}
	
}
