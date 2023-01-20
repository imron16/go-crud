package main

import (
	"go-crud/interface/container"
	"go-crud/interface/server"
	"go-crud/shared/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server.StartApp(container.SetupContainer(database.Connect())).Run()

}
