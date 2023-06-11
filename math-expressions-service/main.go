package main

import (
	"net/http"

	"test-domain.com/math-expressions-service/handlers"
)

func main() {
	config := LoadConfig()

	http.HandleFunc("/evaluate", handlers.EvaluateHandler)
	http.HandleFunc("/validate", handlers.ValidateHandler)
	http.HandleFunc("/errors", handlers.ErrorsHandler)

	http.ListenAndServe(":"+config.ServicePort, nil)
}
