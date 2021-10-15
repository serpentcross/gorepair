package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgtype"
	Sparepart "github.com/serpentcross/gorepair/dto/sparepart"
	"log"
	"net/http"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "postgres"
	driver = "postgres"
	dbname = "fixburo"
)

func setupDB() *sql.DB {
	postgresSqlConnection := driver + "://" + user + ":" + pass + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
	dataBaseConnection, err := sql.Open("postgres", postgresSqlConnection)
	checkError(err)
	return dataBaseConnection
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/spareparts", getSpareparts).Methods("GET")
	router.HandleFunc("/spareparts", createSparepart).Methods("POST")

	fmt.Println("Server listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getSpareparts(w http.ResponseWriter, r *http.Request) {

	db := setupDB()

	printMessage("Getting spareparts...")

	rows, err := db.Query("SELECT * FROM sparepart")

	checkError(err)

	var spareparts []Sparepart

	for rows.Next() {

		var id pgtype.UUID
		var name string
		var available bool
		var artikul string

		err = rows.Scan(&id, &name, &available, &artikul)

		checkError(err)

		spareparts = append(spareparts, Sparepart{Id: id, Name: name, Available: available, Artikul: artikul})

	}

	var response = spareparts

	json.NewEncoder(w).Encode(response)

}

func createSparepart(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	sparepartID := r.FormValue("id")
	sparepartName := r.FormValue("name")

	printMessage(sparepartID)
	printMessage(sparepartName)

}
