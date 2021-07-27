package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jijosg/http-go-mux/pkg/sqllite"
	"github.com/jijosg/http-go-mux/pkg/user"

	"github.com/gorilla/mux"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var conn *sql.DB
	const DB_FILENAME = "./data.db"
	sqllite.CreateNewFile(DB_FILENAME)
	conn = sqllite.OpenConnection(DB_FILENAME)
	log.Println("Connected to data.db")
	defer conn.Close()
	r := mux.NewRouter()

	r.HandleFunc("/deletedb", func(w http.ResponseWriter, r *http.Request) {
		os.Remove(DB_FILENAME)
		fmt.Fprintf(w, "Deleted DB file")
	}).Methods("GET")

	r.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		sqllite.CreateNewUserTable(conn)
		fmt.Fprintf(w, "Created table user")
	}).Methods("GET")

	r.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		keyVal := make(map[string]string)
		json.Unmarshal(body, &keyVal)
		log.Println(keyVal)
		name := keyVal["name"]
		err = sqllite.InsertRow(conn, name, DB_FILENAME)
		if err != nil {
			panic(err.Error())
		}
		log.Printf("Inserted row for %s\n", name)
		fmt.Fprintf(w, "Inserted row for "+name)

	}).Methods("POST")

	r.HandleFunc("/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		rows := mux.Vars(r)
		id := rows["id"]
		err := sqllite.DeleteRow(conn, id, DB_FILENAME)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Printf("Deleted row for %s\n", id)
		fmt.Fprintf(w, "Deleted row for "+id)
	}).Methods("DELETE")

	r.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Listing data from table user")
		insertUserSQL := "select * from user"
		stmt, _ := conn.Query(insertUserSQL)
		defer stmt.Close()
		var users []user.User
		for stmt.Next() {
			var user user.User
			stmt.Scan(&user.Id, &user.Name)
			users = append(users, user)
		}
		result, err := json.MarshalIndent(users, "", " ")
		if err != nil {
			log.Fatalln(err.Error())
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}).Methods("GET")
	http.ListenAndServe(":8000", r)
}
