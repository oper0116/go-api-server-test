package store

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	Repository Repository
}

// Users GET /Index
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// Users GET /users
func (c *Controller) User(w http.ResponseWriter, r *http.Request) {
	products := c.Repository.GetUser() // get User Infomation
	// log.Println(products)
	data, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}