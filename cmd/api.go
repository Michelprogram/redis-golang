package main

import (
	"fmt"
	"handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	var port string = os.Args[1]

	router := mux.NewRouter()

	router.HandleFunc("/redis", handlers.StudentsRedis).Methods("GET")
	router.HandleFunc("/notredis", handlers.StudentsNotRedis).Methods("GET")

	router.NotFoundHandler = handlers.ErrorRoute()

	fmt.Printf("🚀 Lancement de l'api sur le port %s\n", port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Printf("Err start serveur : %v\n", err)
	}
}
