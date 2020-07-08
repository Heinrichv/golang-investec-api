package main

import (
	"fmt"
	"log"
	"net/http"

	controllers "./controllers"
	middleware "./middleware"
	mux "github.com/gorilla/mux"
)

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Use(middleware.JSONResponseMiddleware)
	myRouter.Use(middleware.OAuthMiddleware)

	myRouter = controllers.InitiateAuthRoutes(myRouter)

	fmt.Println("Server listening on (http://localhost:4200)")
	log.Fatal(http.ListenAndServe(":4200", myRouter))
}

func main() {
	handleRequests()
}
