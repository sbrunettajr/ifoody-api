package repository

import (
	"database/sql"

	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

var _ repository.DataManager = dataManager{}

type dataManager struct {
	db *sql.DB
}

func NewDataManager(
	db *sql.DB,
) dataManager {
	return dataManager{
		db: db,
	}
}

func (d dataManager) Begin() (*sql.Tx, error) {
	return d.db.Begin()
}

func (d dataManager) Category() repository.CategoryRepository {
	return newCategoryMySQLRepository(d.db)
}

func (d dataManager) Item() repository.ItemRepository {
	return newItemMySQLRepository(d.db)
}

func (d dataManager) Order() repository.OrderRepository {
	return newOrderMySQLRepository(d.db)
}

func (d dataManager) OrderItem() repository.OrderItemRepository {
	return newOrderItemMySQLRepository(d.db)
}

func (d dataManager) Store() repository.StoreRepository {
	return newStoreMySQLRepository(d.db)
}
