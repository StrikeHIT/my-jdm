package entities

import (
	"encoding/json"
	"net/http"

	"github.com/StrikeHIT/my-jdm/configs"
	"github.com/gorilla/mux"
)

type Brand struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
}

func GetBrand(w http.ResponseWriter, r *http.Request) {
	db := configs.Open()

	brand := []Brand{}
	result := db.Find(&brand)
	json.NewEncoder(w).Encode(result)
}

func GetAllBrands(w http.ResponseWriter, r *http.Request) {
	db := configs.Open()

	brands := []Car{}
	result := db.Find(&brands)
	json.NewEncoder(w).Encode(result)
}

func GetBrandByName(w http.ResponseWriter, r *http.Request) {
	db := configs.Open()

	params := mux.Vars(r)
	name := params["name"]
	brands := []Car{}
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&brands)
	json.NewEncoder(w).Encode(result)
}
