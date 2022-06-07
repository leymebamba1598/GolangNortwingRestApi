package main

import (
	"fmt"

	"github.com/leymebamba1598/GolangNortwingRestApi/database"
)

func main() {

	databaseConnection:=database.InitDB()

	//Logica
	defer databaseConnection.Close() //se ejecuta cuando finaliza la funcion
	fmt.Println("data", databaseConnection)

}


