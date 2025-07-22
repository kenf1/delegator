package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/routes"
	"github.com/kenf1/delegator/src/routes/middleware"
)

func main() {
	serverAddr, err := io.ImportServerAddrWrapper(".env")
	if err != nil {
		log.Fatal(err)
	}

	globalAuthConfig, err := io.ImportAuthConfig()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.HandleEntry)
	mux.Handle("/auth/", http.StripPrefix("/auth", middleware.DefaultCorsMiddleware(
		routes.AuthRoutes(globalAuthConfig), serverAddr, "GET, POST",
	)))
	mux.Handle("/tasks/", http.StripPrefix("/tasks", middleware.DefaultCorsMiddleware(
		routes.TasksRoutes(), serverAddr, "GET, POST, PUT, PATCH, DELETE",
	)))

	fmt.Printf("Server listening to %s:%s\n", serverAddr.Host, serverAddr.Port)
	err1 := http.ListenAndServe(":"+serverAddr.Port, mux)
	if err1 != nil {
		log.Fatal(err1)
	}
}
