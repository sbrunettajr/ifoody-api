package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

var _ repository.ItemRepository = itemMySQLRepository{}

type itemMySQLRepository struct {
	db *sql.DB
}

func newItemMySQLRepository(
	db *sql.DB,
) itemMySQLRepository {
	return itemMySQLRepository{
		db: db,
	}
}

func (r itemMySQLRepository) BulkInsert(context context.Context, items []entity.Item) error {
	query := `
		INSERT
		  INTO tb_item(uuid, code, name, description, price, category_id, store_id)
		VALUES %s
	`

	if len(items) == 0 {
		return nil
	}

	values := strings.TrimSuffix(strings.Repeat("(?, ?, ?, ?, ?, ?, ?), ", len(items)), ",")
	query = fmt.Sprintf(query, values)

	var args []any
	for _, item := range items {
		args = append(
			args,
			item.UUID,
			item.Code,
			item.Name,
			item.Description,
			item.Price,
			item.Category.ID,
			item.Store.ID,
		)
	}

	_, err := r.db.ExecContext(context, query, args)
	if err != nil {
		return err
	}
	return nil
}

func (r itemMySQLRepository) Create(ctx context.Context, item entity.Item) (uint32, error) {
	query := `
		INSERT
		  INTO tb_item(uuid, code, name, description, price, category_id, store_id)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		item.UUID,
		item.Code,
		item.Name,
		item.Description,
		item.Price,
		item.Category.ID,
		item.Store.ID,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint32(id), nil
}

func (r itemMySQLRepository) FindByStoreUUID(context context.Context, storeUUID string) ([]entity.Item, error) {
	query := `
		SELECT ti.id,
		       ti.created_at,
			   ti.updated_at,
			   ti.uuid,
			   ti.code,
			   ti.name,
			   ti.description,
			   ti.price,
			   ti.category_id,
			   ti.store_id 
		  FROM tb_item ti 
		  JOIN tb_store ts ON ts.id = ti.store_id
		 WHERE ti.deleted_at IS NULL
		   AND ts.uuid = ?;
	`

	rows, err := r.db.QueryContext(
		context,
		query,
		storeUUID,
	)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	if err != nil {
		return nil, err
	}

	items, err := parseEntities[entity.Item](rows, r.parseEntity)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r itemMySQLRepository) FindByCategoryUUID(context context.Context, categoryUUID string) ([]entity.Item, error) {
	query := `
		SELECT ti.id,
		       ti.created_at,
			   ti.updated_at,
			   ti.uuid,
			   ti.code,
			   ti.name,
			   ti.description,
			   ti.price,
			   ti.category_id,
			   ti.store_id 
		  FROM tb_item ti 
		  JOIN tb_category tc ON tc.id = ti.category_id AND tc.deleted_at IS NULL
		 WHERE ti.deleted_at IS NULL
		   AND tc.uuid = ?; 
	`

	rows, err := r.db.QueryContext(
		context,
		query,
		categoryUUID,
	)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	if err != nil {
		return nil, err
	}

	items, err := parseEntities[entity.Item](rows, r.parseEntity)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (r itemMySQLRepository) FindByUUID(context context.Context, UUID string) (entity.Item, error) {
	query := `
		SELECT ti.id,
		       ti.created_at,
			   ti.updated_at,
			   ti.uuid,
			   ti.code,
			   ti.name,
			   ti.description,
			   ti.price,
			   ti.category_id,
			   ti.store_id 
		  FROM tb_item ti 
		 WHERE ti.deleted_at IS NULL
		   AND ti.uuid = ?; 
	`

	item, err := findByUUID[entity.Item](context, r.db, query, UUID, r.parseEntity)
	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}

func (r itemMySQLRepository) FindByID(context context.Context, ID uint32) (entity.Item, error) {
	query := `
		SELECT ti.id,
		       ti.created_at,
			   ti.updated_at,
			   ti.uuid,
			   ti.code,
			   ti.name,
			   ti.description,
			   ti.price,
			   ti.category_id,
			   ti.store_id 
		  FROM tb_item ti 
		 WHERE ti.deleted_at IS NULL
		   AND ti.id = ?; 
	`

	item, err := findByID[entity.Item](context, r.db, query, ID, r.parseEntity)
	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}

func (r itemMySQLRepository) Delete(context context.Context, UUID string) error {
	query := `
		UPDATE tb_item ti
		   SET ti.deleted_at = NOW()
		 WHERE ti.uuid = ?;  
	`

	_, err := r.db.ExecContext(
		context,
		query,
		UUID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r itemMySQLRepository) Update(context context.Context, item entity.Item) error {
	query := `
		UPDATE tb_item ti
		   SET ti.code = ?,
		       ti.name = ?,
			   ti.description = ?,
			   ti.price = ?,
			   ti.category_id = ?
		 WHERE ti.deleted_at IS NULL
		   AND ti.uuid = ?; 
	`

	_, err := r.db.ExecContext(
		context,
		query,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r itemMySQLRepository) parseEntity(scan scanner) (entity.Item, error) {
	var item entity.Item
	err := scan.Scan(
		&item.ID,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.UUID,
		&item.Code,
		&item.Name,
		&item.Description,
		&item.Price,
		&item.CategoryID,
		&item.StoreID,
	)
	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}
