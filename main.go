package main

import (
	"fmt"
	"go-api-server-test/router"
	"net/http"

	"github.com/gorilla/handlers"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}


var users = map[string]*User{}

type User struct {
    Id string `json:"id"`
    Name string `json:"name"`
}



func main() {
    fmt.Println("START Server Port 8080")

    // mux := http.NewServeMux();

    router := router.NewRouter();

    allowedOrigins := handlers.AllowedOrigins([]string{"*"}) 
 	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
    http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router))

    // mux.Handle("/api/", apiHandler{})
    // mux.Handle("/users", jsonContentTypeMiddleware(userHandler))
    // mux.Handle("/", jsonContentTypeMiddleware(rootHandler))

    // http.ListenAndServe(":8080", mux)
}
