package config

import (
    "database/sql"
    "log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


var DB *sql.DB

func InitDB() (*sql.DB, error) {
	dsn := "root:Lolasso1012@tcp(127.0.0.1:3306)/kafka"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la conexión a la base de datos: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al verificar a la base de datos: %v", err)
	} else {
		fmt.Println("Conexión a la base de datos exitosa")
	}

	DB = db

	return db, nil
}