package main

import (
	"fmt"
	"net/http"
)

// Handler para o endpoint raiz "/"
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	contentType := r.Header.Get("Content-Type")

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello, World! "+contentType)
		return
	}

	// Retorna erro se o método não for GET
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, "Method not allowed")
}

// Handler para o endpoint "/error"
func errorHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 - Not Found "+contentType)
}

func main() {
	// Define os handlers para cada rota
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/error", errorHandler)

	fmt.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
