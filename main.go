package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/santitee/webservices/database"
	"github.com/santitee/webservices/product"
	"github.com/santitee/webservices/receipt"
	"log"
	"net/http"
)


const basePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(basePath)
	product.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}