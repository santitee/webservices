package main

import (
	"github.com/santitee/webservices/product"
	"log"
	"net/http"
)


const basePath = "/api"

func main() {
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}