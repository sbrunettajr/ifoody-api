package repository

import "database/sql"

type DataManager interface {
	Begin() (*sql.Tx, error)
	Category() CategoryRepository
	Item() ItemRepository
	Order() OrderRepository
	OrderItem() OrderItemRepository
	Store() StoreRepository
}
