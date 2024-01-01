package entity

type OrderItem struct {
	base
	UUID     string // Necessary?
	Quantity int
	Price    float64
	Item     Item
}
