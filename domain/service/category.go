package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

type CategoryService struct {
	dataManager  repository.DataManager
	storeService StoreService
}

func NewCategoryService(
	dataManager repository.DataManager,
	storeService StoreService,
) CategoryService {
	return CategoryService{
		dataManager:  dataManager,
		storeService: storeService,
	}
}

func (s CategoryService) Create(context context.Context, category entity.Category) error {
	category.UUID = uuid.NewString()

	store, err := s.storeService.FindByUUID(context, category.Store.UUID)
	if err != nil {
		return err
	}
	category.Store = store

	_, err = s.dataManager.Category().Create(context, category)
	if err != nil {
		return err
	}
	return nil
}

func (s CategoryService) FindByStoreUUID(context context.Context, storeUUID string) ([]entity.Category, error) {
	categories, err := s.dataManager.Category().FindByStoreUUID(context, storeUUID)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (s CategoryService) FindByUUID(context context.Context, UUID string) (entity.Category, error) {
	category, err := s.dataManager.Category().FindByUUID(context, UUID)
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}
