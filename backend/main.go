package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/StrikeHIT/my-jdm/configs"
	"github.com/StrikeHIT/my-jdm/entities"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MAZDA MEDUSA!")
	fmt.Println("Endpoint Hit")
}

func handleRequests() {
	log.Println("Comece a desenvolver http:/localhost:3001/")
	log.Println("Para interromper o servidor use CONTROL-C.")

	//ROUTES
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/api/car/{id}", entities.GetCar)
	myRouter.HandleFunc("/api/car/name/{name}", entities.GetCarByName)
	myRouter.HandleFunc("/api/cars", entities.GetAllCars)

	myRouter.HandleFunc("/api/brand/{id}", entities.GetBrand)
	myRouter.HandleFunc("/api/brand/name/{name}", entities.GetBrandByName)
	myRouter.HandleFunc("/api/brands", entities.GetAllBrands)

	log.Fatal(http.ListenAndServe(":3001", myRouter))
}

func main() {
	db := configs.Open()
	db.AutoMigrate(&entities.Car{}, &entities.Brand{})

	handleRequests()

	defer configs.Close()
}
