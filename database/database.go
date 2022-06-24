package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	connectionString := "root:didier@tcp(localhost:3306)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Error de conexion")
		panic(err.Error())
	}
	return databaseConnection
}