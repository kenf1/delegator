package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/kenf1/delegator/docs"
	"github.com/kenf1/delegator/src/configs"
	"github.com/kenf1/delegator/src/routes"
	"github.com/kenf1/delegator/src/routes/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Delegator
//	@version		1.0
//	@description	Delegator aka microservices entrypoint
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	kenf1
//	@contact.url	http://www.github.com/kenf1

// @license.name	GNU GPLv3
// @license.url	https://www.gnu.org/licenses/gpl-3.0.en.html
func main() {
	serverAddr, err := configs.ImportServerAddrWrapper(".env")
	if err != nil {
		log.Fatal(err)
	}

	globalAuthConfig, err := configs.ImportAuthConfig()
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
	mux.Handle("/docs/", httpSwagger.WrapHandler)

	fmt.Printf("Server listening to %s:%s\n", serverAddr.Host, serverAddr.Port)
	err1 := http.ListenAndServe(":"+serverAddr.Port, mux)
	if err1 != nil {
		log.Fatal(err1)
	}
}
