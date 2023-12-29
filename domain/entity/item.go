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

func (e Item) IsEqual(item Item) bool {
	return e.Code == item.Code &&
		e.Name == item.Name &&
		e.Description == item.Description &&
		e.Price == item.Price &&
		e.Category.Name == item.Category.Name
}
