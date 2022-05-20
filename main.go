package main

import (
	"ramo/bookstore/mappings"
	"ramo/bookstore/models"
)

func main() {

	// Connect to database
	models.ConnectDatabase()
	mappings.CreateUrlMappings()
	//mappings.Router.Run(":8081")
}
