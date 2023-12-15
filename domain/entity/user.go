package entity

type User struct {
	base
	UUID      string
	FirstName string
	LastName  string
	Email     string
	Password  string // Password Salt
}
