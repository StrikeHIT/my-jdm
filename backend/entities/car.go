package entities

import (
	"encoding/json"
	"net/http"

	"github.com/StrikeHIT/my-jdm/configs"
	"github.com/StrikeHIT/my-jdm/utils"
	"github.com/gorilla/mux"
)

type Car struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	BrandID int    `json:"brand_id"`
	Name    string `json:"name"`
	Year    string `json:"year"`
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w, r)
	db := configs.Open()

	car := []Car{}
	result := db.Find(&car).Value
	json.NewEncoder(w).Encode(result)
}

func GetAllCars(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w, r)
	db := configs.Open()

	cars := []Car{}
	result := db.Find(&cars).Value
	json.NewEncoder(w).Encode(result)
}

func GetCarByName(w http.ResponseWriter, r *http.Request) {
	utils.SetHeaders(w, r)
	db := configs.Open()

	params := mux.Vars(r)
	name := params["name"]
	cars := []Car{}
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&cars).Value
	json.NewEncoder(w).Encode(result)
}
