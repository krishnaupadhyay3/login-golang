package model

type User struct {
	Username string
	Password string
	LoggedIn bool
	Email    string
}

//  response return type struct
type ErrorValue struct {
	Error string
}
type Success struct {
	Status string
}

// template struct

type ActiveToggle struct {
	Login  string
	Signup string
}
