package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productJSON := `[
		{
			"productId":1,
			"manufacturer":"Johns-Jenkins",
			"sku":"p5z343vds",
			"upc":"939581000000",
			"pricePerUnit":"497.45",
			"quantityOnHand":3703,
			"productName":"stricky note"
		},
		{
			"productId":2,
			"manufacturer":"Test",
			"sku":"TestSku",
			"upc":"TestUpc",
			"pricePerUnit":"12.99",
			"quantityOnHand":9500,
			"productName":"ProductTest"
		},
		{
			"productId":3,
			"manufacturer":"Swaniawski, Bartoletti and Bruen",
			"sku":"q0L657ys7",
			"upc":"111730000000",
			"pricePerUnit":"439.26",
			"quantityOnHand":5905,
			"productName":"lamp shade"
		}
	]`
	err := json.Unmarshal([]byte(productJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func productsHandler(w http.ResponseWriter, r * http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)

	}
}

func main() {

	http.HandleFunc("/products", productsHandler)

	http.ListenAndServe(":5000", nil)
}
