package main

import (
	"encoding/json"
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

	// - repository
	productsMap := repository.NewProductMapRepository(products, len(products))
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

	router.Route("/products", func(r chi.Router) {
		// define routes
		r.Get("/", hd.GetAllProducts())
		r.Post("/", hd.SaveProduct())
		r.Get("/{id}", hd.GetProduct())
		r.Put("/{id}", hd.UpdateProduct())
		r.Patch("/{id}", hd.PartialUpdateProduct())
		r.Delete("/{id}", hd.DeleteProduct())
		r.Get("/search", hd.GetProductsByPrice())
	})

	// start server
	http.ListenAndServe(":8080", router)

}
