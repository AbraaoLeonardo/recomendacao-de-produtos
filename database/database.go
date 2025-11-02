package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(user, password, dbname, host, port string) {
	var err error
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, dbname, host, port)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão (sql.Open): %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar/pingar o banco de dados: %v", err)
	}

	log.Println("Pool de Conexões com o PostgreSQL estabelecido com sucesso!")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Pool de Conexões com o PostgreSQL fechado.")
	}
}
