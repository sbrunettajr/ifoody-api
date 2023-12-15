package repository

import (
	"context"
	"database/sql"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

var _ repository.CategoryRepository = categoryMySQLRepository{}

type categoryMySQLRepository struct {
	db *sql.DB
}

func newCategoryMySQLRepository(
	db *sql.DB,
) categoryMySQLRepository {
	return categoryMySQLRepository{
		db: db,
	}
}

func (r categoryMySQLRepository) Create(context context.Context, category entity.Category) error {
	query := `
		INSERT 
		  INTO tb_category(uuid, name, store_id)
		VALUES (?, ?, ?);  
	`

	_, err := r.db.ExecContext(
		context,
		query,
		category.UUID,
		category.Name,
		category.Store.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r categoryMySQLRepository) FindByStoreUUID(context context.Context, storeUUID string) ([]entity.Category, error) {
	query := `
		SELECT tc.id,
		       tc.created_at,
			   tc.updated_at,
			   tc.uuid,
			   tc.name,
			   tc.store_id
		  FROM tb_category tc 
		  JOIN tb_store ts ON ts.id = tc.store_id AND ts.deleted_at IS NULL
		 WHERE tc.deleted_at IS NULL
		   AND ts.uuid = ?; 
	`

	rows, err := r.db.QueryContext(
		context,
		query,
		storeUUID,
	)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	categories := make([]entity.Category, 0)
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(
			&category.ID,
			&category.CreatedAt,
			&category.UpdatedAt,
			&category.UUID,
			&category.Name,
			&category.StoreID,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r categoryMySQLRepository) FindByUUID(context context.Context, UUID string) (entity.Category, error) {
	query := `
		SELECT tc.id,
		       tc.created_at,
			   tc.updated_at,
			   tc.uuid,
			   tc.name,
			   tc.store_id
		  FROM tb_category tc 
		 WHERE tc.deleted_at IS NULL
		   AND tc.uuid = ?; 
	`

	row := r.db.QueryRowContext(
		context,
		query,
		UUID,
	)
	if row.Err() != nil {
		return entity.Category{}, row.Err()
	}

	var category entity.Category
	err := row.Scan(
		&category.ID,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.UUID,
		&category.Name,
		&category.StoreID,
	)
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}
