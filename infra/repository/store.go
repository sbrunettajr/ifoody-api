package repository

import (
	"context"
	"database/sql"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

var _ repository.StoreRepository = storeMySQLRepository{}

type storeMySQLRepository struct {
	db *sql.DB
}

func newStoreMySQLRepository(
	db *sql.DB,
) storeMySQLRepository {
	return storeMySQLRepository{
		db: db,
	}
}

func (r storeMySQLRepository) Create(context context.Context, store entity.Store) error {
	query := `
		INSERT 
		  INTO tb_store(uuid, fantasy_name, corporate_name)
		VALUES (?, ?, ?);  
	`

	_, err := r.db.ExecContext(
		context,
		query,
		store.UUID,
		store.FantasyName,
		store.CorporateName,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r storeMySQLRepository) FindAll(context context.Context) ([]entity.Store, error) {
	query := `
		SELECT ts.id,
		       ts.created_at,
			   ts.updated_at,
			   ts.uuid,
			   ts.fantasy_name,
			   ts.corporate_name
		  FROM tb_store ts 
		 WHERE ts.deleted_at IS NULL; 
	`

	rows, err := r.db.QueryContext(
		context,
		query,
	)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}

	stores := make([]entity.Store, 0)
	for rows.Next() {
		var store entity.Store
		err := rows.Scan(
			&store.ID,
			&store.CreatedAt,
			&store.UpdatedAt,
			&store.UUID,
			&store.FantasyName,
			&store.CorporateName,
		)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}
	return stores, nil
}

func (r storeMySQLRepository) FindByFantasyName(context context.Context, FantasyName string) (entity.Store, error) {
	query := `
		SELECT ts.id,
		       ts.created_at,
			   ts.updated_at,
			   ts.uuid,
			   ts.fantasy_name,
			   ts.corporate_name
		  FROM tb_store ts 
		 WHERE ts.deleted_at IS NULL
		   AND ts.fantasy_name = ?; 
	`

	row := r.db.QueryRowContext(
		context,
		query,
		FantasyName,
	)
	if row.Err() != nil {
		return entity.Store{}, row.Err()
	}

	var store entity.Store
	err := row.Scan(
		&store.ID,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.UUID,
		&store.FantasyName,
		&store.CorporateName,
	)
	if err != nil {
		return entity.Store{}, err
	}
	return store, nil
}

func (r storeMySQLRepository) FindByUUID(context context.Context, UUID string) (entity.Store, error) {
	query := `
		SELECT ts.id,
		       ts.created_at,
			   ts.updated_at,
			   ts.uuid,
			   ts.fantasy_name,
			   ts.corporate_name
		  FROM tb_store ts 
		 WHERE ts.deleted_at IS NULL
		   AND ts.uuid = ?; 
	`

	row := r.db.QueryRowContext(
		context,
		query,
		UUID,
	)
	if row.Err() != nil {
		return entity.Store{}, row.Err()
	}

	var store entity.Store
	err := row.Scan(
		&store.ID,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.UUID,
		&store.FantasyName,
		&store.CorporateName,
	)
	if err != nil {
		return entity.Store{}, err
	}
	return store, nil
}
