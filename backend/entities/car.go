package entities

import (
	"encoding/json"
	"net/http"

	"github.com/StrikeHIT/my-jdm/configs"
	"github.com/gorilla/mux"
)

type Car struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	BrandID int    `json:"brand_id"`
	Name    string `json:"name"`
	Year    string `json:"year"`
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	db := configs.Open()

	car := []Car{}
	result := db.Find(&car)
	json.NewEncoder(w).Encode(result)
}

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	db := configs.Open()

	cars := []Car{}
	result := db.Find(&cars)
	json.NewEncoder(w).Encode(result)
}

func GetCarByName(w http.ResponseWriter, r *http.Request) {
	db := configs.Open()

	params := mux.Vars(r)
	name := params["name"]
	cars := []Car{}
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&cars)
	json.NewEncoder(w).Encode(result)
}
