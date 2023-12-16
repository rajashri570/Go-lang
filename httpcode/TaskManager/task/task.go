package task

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time.Time"
)

var DB *gorm.DB
var err error

const DNS = "root:root123@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type Task struct {
	gorm.Model
	//Id       int       `json:"id"`
	Status   int       `json:"status"`
	Priority int       `json:"priority"`
	Deadline time.Time `json:"deadline"`
	isvalid  bool      `json: "isvalid"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Task{})
}

func View_tasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tasks []Task
	DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func Get_task(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var task Task
	DB.First(&task, params["id"])
	json.NewEncoder(w).Encode(task)

}
func Create_task(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	DB.Create(&task)
	json.NewEncoder(w).Encode(task)

} /*
func delete_task(w http.ResponseWriter, r *http.Request) {

}
*/
