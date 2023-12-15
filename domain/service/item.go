package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

type ItemService struct {
	categoryService CategoryService
	dataManager     repository.DataManager
	storeService    StoreService
}

func NewItemService(
	categoryService CategoryService,
	dataManager repository.DataManager,
	storeService StoreService,
) ItemService {
	return ItemService{
		categoryService: categoryService,
		dataManager:     dataManager,
		storeService:    storeService,
	}
}

func (s ItemService) Create(context context.Context, item entity.Item) error {
	category, err := s.categoryService.FindByUUID(context, item.Category.UUID)
	if err != nil {
		return err
	}
	item.Category = category

	store, err := s.storeService.FindByUUID(context, item.Store.UUID)
	if err != nil {
		return err
	}
	item.Store = store

	item.UUID = uuid.NewString()

	err = s.dataManager.Item().Create(context, item)
	if err != nil {
		return err
	}
	return nil
}

func (s ItemService) FindAll(context context.Context) ([]entity.Item, error) {
	items, err := s.dataManager.Item().FindAll(context)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s ItemService) FindByUUID(context context.Context, UUID string) (entity.Item, error) {
	item, err := s.dataManager.Item().FindByUUID(context, UUID)
	if err != nil {
		return entity.Item{}, err
	}
	return item, nil
}

func (s ItemService) FindByCategoryUUID(context context.Context, categoryUUID string) ([]entity.Item, error) {
	items, err := s.dataManager.Item().FindByCategoryUUID(context, categoryUUID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s ItemService) Delete(context context.Context, UUID string) error {
	err := s.dataManager.Item().Delete(context, UUID)
	if err != nil {
		return err
	}
	return nil
}

func (s ItemService) Update(context context.Context, item entity.Item) error {
	err := s.dataManager.Item().Update(context, item)
	if err != nil {
		return err
	}
	return nil
}
