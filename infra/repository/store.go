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

func (r storeMySQLRepository) Create(context context.Context, store entity.Store) (uint32, error) {
	query := `
		INSERT 
		  INTO tb_store(uuid, fantasy_name, corporate_name, cnpj)
		VALUES (?, ?, ?, ?);
	`

	result, err := r.db.ExecContext(
		context,
		query,
		store.UUID,
		store.FantasyName,
		store.CorporateName,
		store.CNPJ,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint32(id), err
}

func (r storeMySQLRepository) FindAll(context context.Context) ([]entity.Store, error) {
	query := `
		SELECT ts.id,
		       ts.created_at,
			   ts.updated_at,
			   ts.uuid,
			   ts.fantasy_name,
			   ts.corporate_name,
			   ts.cnpj
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

	stores, err := parseEntities[entity.Store](rows, r.parseEntity)
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (r storeMySQLRepository) FindByCNPJ(context context.Context, CNPJ string) (entity.Store, error) {
	query := `
		SELECT ts.id,
		       ts.created_at,
			   ts.updated_at,
			   ts.uuid,
			   ts.fantasy_name,
			   ts.corporate_name,
			   ts.cnpj
		  FROM tb_store ts 
		 WHERE ts.deleted_at IS NULL
		   AND ts.cnpj = ?; 
	`

	row := r.db.QueryRowContext(
		context,
		query,
		CNPJ,
	)
	if row.Err() != nil {
		return entity.Store{}, row.Err()
	}

	store, err := r.parseEntity(row)
	if err != nil {
		return entity.Store{}, err
	}
	return store, nil
}

func (r storeMySQLRepository) FindByID(context context.Context, ID uint32) (entity.Store, error) {
	query := `
		SELECT ts.id,
		       ts.created_at,
			   ts.updated_at,
			   ts.uuid,
			   ts.fantasy_name,
			   ts.corporate_name,
			   ts.cnpj
		  FROM tb_store ts 
		 WHERE ts.deleted_at IS NULL
		   AND ts.id = ?; 
	`

	store, err := findByID[entity.Store](context, r.db, query, ID, r.parseEntity)
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
			   ts.corporate_name,
			   ts.cnpj
		  FROM tb_store ts 
		 WHERE ts.deleted_at IS NULL
		   AND ts.uuid = ?; 
	`

	store, err := findByUUID[entity.Store](context, r.db, query, UUID, r.parseEntity)
	if err != nil {
		return entity.Store{}, err
	}
	return store, nil
}

func (r storeMySQLRepository) parseEntity(scan scanner) (entity.Store, error) {
	var store entity.Store
	err := scan.Scan(
		&store.ID,
		&store.CreatedAt,
		&store.UpdatedAt,
		&store.UUID,
		&store.FantasyName,
		&store.CorporateName,
		&store.CNPJ,
	)
	if err != nil {
		return entity.Store{}, err
	}
	return store, nil
}
