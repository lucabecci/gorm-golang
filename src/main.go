package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lucabecci/gorm-golang/src/controllers"
	"github.com/lucabecci/gorm-golang/src/database"
)

func main() {
	//Enviroment Configuration
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	port, _ := os.LookupEnv("PORT")

	//Database Call
	db := database.GetConnection()
	if db.Success == false {
		log.Panic("DB is not connected", db.Error)
	}
	fmt.Println("DB is connected")

	//Server and Router
	r := mux.NewRouter()
	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	r.HandleFunc("/api/task", controllers.CreateTask).Methods("POST")
	r.HandleFunc("/api/task/{id}", controllers.GetTask).Methods("GET")
	r.HandleFunc("/api/task", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/api/task/{id}", controllers.UpdateTask).Methods("PUT")
	r.HandleFunc("/api/task/{id}", controllers.DeleteTask).Methods("DELETE")

	fmt.Println("Server running at port:", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
