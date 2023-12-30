package entity

type OrderItem struct {
	base
	UUID     string // Necessary?
	Quantity int
	Item     Item
}
