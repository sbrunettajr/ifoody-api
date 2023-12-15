package entity

type Order struct {
	base
	UUID   string
	User   User
	Status string // Create status constants
}
