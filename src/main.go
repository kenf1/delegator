package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/models"
	"github.com/kenf1/delegator/src/routes"
)

func main() {
	serverAddr, err := io.ImportServerAddrWrapper("../.env")
	if err != nil {
		log.Fatal(err)
	}

	globalAuthConfig := models.AuthConfig{
		SecretKey: []byte("42069"),
		Issuer:    "me",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.HandleEntry)

	//jwt token
	mux.HandleFunc("POST /auth/create", routes.GenerateJWT(globalAuthConfig))
	mux.HandleFunc("GET /auth/uncreate/{token}", routes.DeconstructJWT(globalAuthConfig))

	//tasks: use in-memory database
	mux.HandleFunc("GET /tasks", routes.ReadAllTasks)
	mux.HandleFunc("GET /tasks/{id}", routes.ReadSingleTask)
	mux.HandleFunc("POST /tasks", routes.CreateTask)
	mux.HandleFunc("DELETE /tasks/{id}", routes.DeleteTask)
	mux.HandleFunc("PUT /tasks", routes.PutTask)
	mux.HandleFunc("PATCH /tasks", routes.PatchTask)

	fmt.Printf("Server listening to %s:%s\n", serverAddr.Host, serverAddr.Port)
	err1 := http.ListenAndServe(":"+serverAddr.Port, mux)
	if err1 != nil {
		log.Fatal(err1)
	}
}
