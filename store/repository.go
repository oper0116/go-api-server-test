package store

import "go-api-server-test/models"

type Repository struct{}

func (r Repository) GetUser() models.User {
	obj := &models.User{ID: "1", NAME: "handonghee", EMAIL: "oper0116@gmail.com"}
	return *obj
}