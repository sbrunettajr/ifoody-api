package entity

type Category struct {
	base
	UUID    string
	Name    string
	StoreID uint32
	Store   Store
}
