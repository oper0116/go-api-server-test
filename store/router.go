package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route {
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route {
		"User",
		"GET",
		"/user",
		controller.User,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes { 
			var handler http.Handler
			log.Println(route.Name)
			handler = route.HandlerFunc
			router.
			 Methods(route.Method).
			 Path(route.Pattern).
			 Name(route.Name).
			 Handler(handler)

	}
	return router
}


// func ContentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			w.Header().Set("Content-Type", "application/json")
// 			next.ServeHTTP(w, r);
// 	})
// }