package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	data := []*product{
		&product{ID: 1, Name: "p1"},
		&product{ID: 2, Name: "p2"},
		&product{ID: 3, Name: "p3"},
	}

	router := mux.NewRouter()

	// GET /products
	router.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			log.Print("GET /products")
			d, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("json marshal: %v", err)
				return
			}
			log.Printf("DATA: %s", string(d))
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			w.Write(d)
		}
	}).Methods(http.MethodGet)

	// POST /product
	router.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			w.Header().Set("Access-Control-Allow-Methods", "POST")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			log.Print("OPTION /product")
		}
		if r.Method == http.MethodPost {
			body := make([]byte, r.ContentLength)
			r.Body.Read(body)

			var p product
			if err := json.Unmarshal(body, &p); err != nil {
				log.Printf("json unmarshal: %v", err)
			}

			data = append(data, &p)
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			log.Printf("POST /product | %v\n", p)
		}
	}).Methods(http.MethodPost, http.MethodOptions)

	server := http.Server{
		Addr:    "localhost:9000",
		Handler: router,
	}

	log.Printf("Server Running: %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
