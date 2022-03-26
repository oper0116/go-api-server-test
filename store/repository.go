package store

type Repository struct{}

func (r Repository) GetUser() User {
	obj := &User{ID: "1", NAME: "handonghee", EMAIL: "oper0116@gmail.com"}
	return *obj
}