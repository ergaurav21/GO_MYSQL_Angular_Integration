package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"webservices/database"
	"webservices/product"
	"webservices/receipt"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	product.SetUP(basePath)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
