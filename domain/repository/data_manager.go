package repository

type DataManager interface {
	Category() CategoryRepository
	Item() ItemRepository
	Store() StoreRepository
}
