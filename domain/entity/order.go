package entity

type Order struct {
	base
	UUID       string
	Status     string
	OrderItems []OrderItem
	Store      Store
}
