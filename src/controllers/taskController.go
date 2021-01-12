package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/lucabecci/gorm-golang/src/helpers"
	"github.com/lucabecci/gorm-golang/src/models"
)

//Data is a struct for the return in controllers
type Data struct {
	Success bool          `json:"success"`
	Data    []models.Task `json:"data"`
	Errors  []string      `json:"errors"`
}

//CreateTask is a function for the endpoint create task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	taskBody, s := helpers.DecodeBody(r)
	if s != true {
		http.Error(w, "Error for decode your body, send a valid data", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}

	taskBody.Title = strings.TrimSpace(taskBody.Title)
	taskBody.Description = strings.TrimSpace(taskBody.Description)
	taskBody.Priority = strings.TrimSpace(taskBody.Priority)

	if !helpers.ValidateInformation(taskBody.Title, taskBody.Description, taskBody.Priority) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid Information, please send a valid information")

		json, _ := json.Marshal(data)

		w.Header().Set("content-type", "application/json")

		w.WriteHeader(http.StatusBadRequest)

		w.Write(json)
		return
	}
	task, success := models.InsertTask(taskBody.Title, taskBody.Description, taskBody.Priority)
	if success == false {
		data.Errors = append(data.Errors, "Error to create your task")
		return
	}

	data.Success = true
	data.Data = append(data.Data, task)

	json, _ := json.Marshal(data)

	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//GetTask is a function for get a task by id of the db
func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data Data
	var task models.Task
	var success bool

	task, success = models.GetTask(id)

	if success == false {
		data.Success = false
		data.Errors = append(data.Errors, "task not found")

		json, _ := json.Marshal(data)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	data.Success = true

	data.Data = append(data.Data, task)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//GetTasks is a function for get all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	var data Data
	tasks, success := models.GetTasks()

	if success != true {
		data.Success = false
		data.Errors = append(data.Errors, "You dont have a tasks")

		json, _ := json.Marshal(data)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}

	data.Success = true

	data.Data = append(data.Data, tasks...)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//UpdateTask is a function for update a task
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskBody, s := helpers.DecodeBody(r)
	vars := mux.Vars(r)
	id := vars["id"]

	if s != true {
		http.Error(w, "Error for decode your body, send a valid data", http.StatusBadRequest)
		return
	}

	var data Data = Data{Errors: make([]string, 0)}

	taskBody.Title = strings.TrimSpace(taskBody.Title)
	taskBody.Description = strings.TrimSpace(taskBody.Description)
	taskBody.Priority = strings.TrimSpace(taskBody.Priority)

	if !helpers.ValidateInformation(taskBody.Title, taskBody.Description, taskBody.Priority) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid Information, please send a valid information")

		json, _ := json.Marshal(data)

		w.Header().Set("content-type", "application/json")

		w.WriteHeader(http.StatusBadRequest)

		w.Write(json)
		return
	}
	task, success := models.UpdateTask(id, taskBody.Title, taskBody.Description, taskBody.Priority)
	if success == false {
		data.Errors = append(data.Errors, "Error to update your task, please check your ID")
		json, _ := json.Marshal(data)

		w.Header().Set("content-type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	data.Success = true
	data.Data = append(data.Data, task)

	json, _ := json.Marshal(data)

	w.Header().Set("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

//DeleteTask is a function for delete a task of the db
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var data Data
	var task models.Task
	var success bool

	task, success = models.DeleteTask(id)

	if success == false {
		data.Success = false
		data.Errors = append(data.Errors, "Error to delete your Task, ID not found")

		json, _ := json.Marshal(data)

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		return
	}
	data.Success = true

	data.Data = append(data.Data, task)

	json, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
