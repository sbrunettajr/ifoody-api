package entity

type Item struct {
	base
	UUID        string
	Code        string
	Name        string
	Description string
	Price       float64 // Fix: monetary value!
	CategoryID  uint32
	StoreID     uint32
	Category    Category
	Store       Store
}
