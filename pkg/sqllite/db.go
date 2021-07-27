package sqllite

import (
	"database/sql"
	"log"
	"os"
)

var createUserTableSQL = `CREATE TABLE user (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
	"name" TEXT
  );`

func CreateNewFile(name string) {

	log.Println("Creating data.db...")

	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err := os.Create(name) // Create SQLite file
		if err != nil {
			log.Println("Could not create file")
		}
		file.Close()
		log.Println("data.db created")
	} else {
		log.Println("File already exists")
	}
}

func OpenConnection(name string) *sql.DB {
	conn, err := sql.Open("sqlite3", name)
	if err != nil {
		log.Fatalln(err.Error())
	}

	return conn
}

func CreateNewUserTable(conn *sql.DB) {
	stmt, err := conn.Prepare(createUserTableSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	stmt.Exec()
}

func InsertRow(conn *sql.DB, name string, dbFileName string) error {
	log.Println("Inserting row into file")
	insertUserSQL := `INSERT INTO user (name) values(?);`
	stmt, err := conn.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRow(conn *sql.DB, id string, dbFileName string) error {
	insertUserSQL := `DELETE FROM user where id = ?;`
	stmt, err := conn.Prepare(insertUserSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
