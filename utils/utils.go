package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type H map[string]interface{}

type ServerResponse struct {
	Status     bool             `json:"status"`
	Message    string           `json:"message"`
	Data       *json.RawMessage `json:"data,omitempty"`
	StatusCode int              `json:"statusCode"`
}

func ShouldBindJSON(r *http.Request, body interface{}) error {
	req := json.NewDecoder(r.Body)
	req.DisallowUnknownFields()

	err := req.Decode(&body)
	if err != nil {

	}

	return nil
}

func JSON(w http.ResponseWriter, statusCode int, data H) {
	res, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

func EnvString(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return ""
	}

	return value
}
