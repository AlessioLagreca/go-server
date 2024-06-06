package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {	

	log.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

	// definire il percorso del file
	path := r.URL.Path

	// se il percorso è index.html o / allora restituisci index.html
	if path == "/" || path == "index.html" {
		path = "index.html"
	}

	// restituisci il file
	http.ServeFile(w, r, path)


	// log all headers
	for key, values := range r.Header {
		for _, value := range values {
			log.Printf("Header: %s: %s\n", key, value)
		}
	}
	// crea la risposta
	response := fmt.Sprintf("Requested path: %s\r\n", r.URL.Path)
	w.Write([]byte(response))
	
	 // Crea lo status code
	 w.WriteHeader(http.StatusOK)
}

// FUNZIONE MAIN

func main() {
	http.HandleFunc("/", handler)
	
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}

