package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
)

type StoreService struct {
	dataManager    repository.DataManager
	metricsService MetricsService
}

func NewStoreService(
	dataManager repository.DataManager,
	metricsService MetricsService,
) StoreService {
	return StoreService{
		dataManager:    dataManager,
		metricsService: metricsService,
	}
}

func (s StoreService) Create(context context.Context, store entity.Store) error {
	store.UUID = uuid.NewString()

	_, err := s.dataManager.Store().Create(context, store)
	if err != nil {
		return err
	}
	return nil
}

func (s StoreService) FindAll(context context.Context) ([]entity.Store, error) {
	stores, err := s.dataManager.Store().FindAll(context)
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (s StoreService) FindByUUID(context context.Context, UUID string) (entity.Store, error) {
	store, err := s.dataManager.Store().FindByUUID(context, UUID)
	if err != nil {
		return entity.Store{}, err
	}
	return store, nil
}
