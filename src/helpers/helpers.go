package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/lucabecci/gorm-golang/src/models"
)

//DecodeBody is a function for check the body in the req
func DecodeBody(r *http.Request) (models.Task, bool) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		return models.Task{}, false
	}
	return task, true
}

//ValidateInformation is a func for check te longitude of the camps
func ValidateInformation(title string, description string, priority string) bool {
	if len(title) < 1 || len(description) < 1 || len(priority) < 1 {
		return false
	}
	return true
}
