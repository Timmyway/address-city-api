package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Data struct {
	ID string `json:id`	
	CodeCommune string `json:code_commune_insee` 
	NomCommune string `json:nom_commune`
	Zipcode string `json:code_postal`
	LibelleAcheminement string `json:libelle_acheminement`
}

func allDatas(w http.ResponseWriter, r *http.Request) {
	db := db_connect()	
	// records := db_get(*db)
	params := mux.Vars(r)
	records := db_filter(*db, params["code_commune_insee"])

	defer db.Close()	
	// datas := Data{Zipcode: records.Zipcode, CodeCommune: records.CodeCommune, NomCommune: records.NomCommune, LibelleAcheminement: records.LibelleAcheminement}
	fmt.Println("Endpoint hit: all datas")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page endpoint hit")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/datas/{cp:[1-9]?[0-9]{0,4}}", allDatas).Methods("GET")	
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func db_connect() (*sql.DB) {
	// Connexion to MySQL database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lp_datas")
	if err != nil {
		panic(err.Error())
	}	
	fmt.Println("Successfully connected to MYSQL database")
	return db
}
func db_insert(db sql.DB) {
	// Insert into database
	insert, err := db.Query("INSERT INTO cp_city(code_commune_insee, nom_commune, code_postal, libelle_acheminement) VALUES('68359', 'WATTWILLER', '68700', 'ABAUCOURT HAUTECOURT')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()	
}

func db_get(db sql.DB) ([]Data) {
	var results []Data
	rows, err := db.Query("SELECT * FROM cp_city")
	if err != nil {
		panic(err.Error())
	}
	var data Data
	for rows.Next() {		

		err = rows.Scan(&data.ID, &data.Zipcode, &data.CodeCommune, &data.NomCommune, &data.LibelleAcheminement)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(data)
		results = append(results, data)
	}
	return results
}

func db_filter(db sql.DB, pattern string) ([]Data) {
	var results []Data
	fmt.Println("Pattern", pattern)
	rows, err := db.Query("SELECT * FROM cp_city WHERE code_commune_insee LIKE '%" + pattern + "%'")
	if err != nil {
		panic(err.Error())
	}
	var data Data
	for rows.Next() {		

		err = rows.Scan(&data.ID, &data.CodeCommune, &data.NomCommune, &data.Zipcode, &data.LibelleAcheminement)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(data)
		results = append(results, data)
	}
	return results
}

func main() {	
	// db := db_connect()
	// db_insert(*db)

	// db_get(*db)

	// defer db.Close()
		
	handleRequests()	
}