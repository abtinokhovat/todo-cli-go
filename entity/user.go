package entity

type User struct {
	Id       uint
	Email    string
	Password string
}

func NewUser(id int, email, password string) *User {
	hashed := hash(password)
	return &User{
		Id:       uint(id),
		Email:    email,
		Password: hashed,
	}
}

func hash(str string) string {
	// TODO: implement hashing
	return str
}
