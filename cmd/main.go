package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/lauramafoMeli/go-web/internal"
	"github.com/lauramafoMeli/go-web/internal/handlers"
	"github.com/lauramafoMeli/go-web/internal/repository"
	"github.com/lauramafoMeli/go-web/internal/service"
)

func main() {
	// Open our jsonFile
	file, err := os.Open("../docs/products.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// we initialize our products array
	var products []internal.Product

	// we unmarshal (convert bytes to struct) our byteArray which contains our jsonFile's content into 'products' which we defined above
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		panic(err)
	}
	productsMap := repository.NewProductMapRepository(products, len(products))
	fmt.Print(productsMap.LastID)

	// - service
	sv := service.NewProductDefault(productsMap)
	// - handler
	hd := handlers.NewDefaultProduct(sv)

	// create a router with chi
	router := chi.NewRouter()

	// define ping route with 200 response
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	// define products route
	router.Get("/products", hd.GetAllProducts())

	router.Post("/products", hd.SaveProduct())

	// define products/:id route
	router.Get("/products/{id}", hd.GetProduct())
	router.Put("/products/{id}", hd.UpdateProduct())
	router.Patch("/products/{id}", hd.PartialUpdateProduct())

	// define products/search route
	router.Get("/products/search", hd.GetProductsByPrice())

	// start server
	http.ListenAndServe(":8080", router)

}
