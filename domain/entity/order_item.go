package entity

type OrderItem struct {
	base
	UUID      string // Necessary?
	Quantity  int
	UnitPrice float64 // Monetary value
}
