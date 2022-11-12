package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(router *mux.Router){

	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")

}

func main() {
	// Load Configurtion from config.json using Viper
	LoadAppConfig()

	// Initialize the router
	database.Connect(AppConfig.ConnectionString)
    database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	//log.Fatal(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}




//  Source blog ->>  https://codewithmukesh.com/blog/implementing-crud-in-golang-rest-api/