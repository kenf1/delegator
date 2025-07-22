package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kenf1/delegator/src/io"
	"github.com/kenf1/delegator/src/routes"
)

func main() {
	serverAddr, err := io.ImportServerAddrWrapper("../.env")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.HandleEntry)

	mux.HandleFunc("GET /tasks", routes.ReadAllTasks)
	mux.HandleFunc("GET /tasks/{id}", routes.ReadSingleTask)
	mux.HandleFunc("POST /tasks", routes.CreateTask)
	mux.HandleFunc("DELETE /tasks/{id}", routes.DeleteTask)
	mux.HandleFunc("PUT /tasks", routes.PutTask)
	mux.HandleFunc("PATCH /tasks", routes.PatchTask)

	mux.HandleFunc("GET /auth/{value}", routes.GenerateJWT)
	mux.HandleFunc("GET /auth1/{value}", routes.DeconstructJWT)

	fmt.Printf("Server listening to %s:%s\n", serverAddr.Host, serverAddr.Port)
	err1 := http.ListenAndServe(":"+serverAddr.Port, mux)
	if err1 != nil {
		log.Fatal(err1)
	}
}
