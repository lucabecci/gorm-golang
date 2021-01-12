package controllers

import "github.com/lucabecci/gorm-golang/src/models"

type Data struct {
	Success bool          `json: "success"`
	Data    []models.Task `json: "data"`
	Errors  []string      `json: "errors"`
}
