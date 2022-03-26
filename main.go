package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var users = map[string]*User{}

type User struct {
    Id string `json:"id"`
    Name string `json:"name"`
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        next.ServeHTTP(w, r);
    })
}

func main() {
    fmt.Println("START Server Port 8080")

    mux := http.NewServeMux();
    
    rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })

    userHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case http.MethodGet:
                json.NewEncoder(w).Encode((users))
            case http.MethodPost:
                var user User
                json.NewDecoder(r.Body).Decode(&user)
                users[user.Id] = &user

                json.NewEncoder(w).Encode(user)
                fmt.Printf("%+v", user)
        }
    })

    mux.Handle("/", jsonContentTypeMiddleware(rootHandler))
    mux.Handle("/users", jsonContentTypeMiddleware(userHandler))

    http.ListenAndServe(":8080", mux)
}
