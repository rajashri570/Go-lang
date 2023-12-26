package main

import (
	"MyProject/task"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/tasks", task.View_tasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", task.Get_task).Methods("GET")
	r.HandleFunc("/tasks", task.Create_task).Methods("POST")
	//r.HandleFunc("/tasks/{id}", update_task).Methods("PUT")
	//r.HandleFunc("/tasks/{id}", delete_task).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))

}
func main() {

	InitialMigration()
	initializeRouter()
}
