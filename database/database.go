package database

import (
	"log"
	"net/http"
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {
	connStr := "user=yourusername dbname=yourdbname sslmode=disable password=yourpassword host=yourhost port=yourport"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	log.Println("Successfully connected to the database")
	return db
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Println("Error closing the database connection: ", err)
	} else {
		log.Println("Database connection closed successfully")
	}
}

func QueryDB(query string)(*sql.Rows, error) {
	db := ConnectToDB()
	defer CloseDB(db)

	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query: ", err)
		return nil, err
	}

	return rows, nil
}

func HandleDBError(w http.ResponseWriter, err error) {
	http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
}

func UpdateDB(query string) error {
	db := ConnectToDB()
	defer CloseDB(db)

	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error executing update: ", err)
		return err
	}

	return nil
}