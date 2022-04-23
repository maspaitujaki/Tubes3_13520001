package main

import (
	"dna-matching-api/controllers"
	"dna-matching-api/database"
	"dna-matching-api/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8080")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8080", router))

}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/penyakit/create", controllers.CreatePenyakit).Methods("POST")
	router.HandleFunc("/penyakit/get", controllers.GetAllPenyakit).Methods("GET")
	router.HandleFunc("/penyakit/get/{nama}", controllers.GetPenyakitByNama).Methods("GET")
	router.HandleFunc("/penyakit/update/{nama}", controllers.UpdatePenyakitByNama).Methods("PUT")
	router.HandleFunc("/penyakit/delete/{nama}", controllers.DeletePenyakitByNama).Methods("DELETE")

	router.HandleFunc("/pemeriksaan/get", controllers.GetAllPemeriksaan).Methods("GET")
	router.HandleFunc("/pemeriksaan/create", controllers.CreatePemeriksaan).Methods("POST")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "@kpop1511",
			DB:         "tubes3_13520001",
		}
	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.MigratePenyakit(&entity.Penyakit{})
	database.MigratePemeriksaan(&entity.Pemeriksaan{})
}
