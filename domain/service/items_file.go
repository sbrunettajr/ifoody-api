package service

import (
	"context"
	"io"
	"strconv"

	"github.com/google/uuid"
	"github.com/sbrunettajr/ifoody-api/domain/constant"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/sbrunettajr/ifoody-api/domain/repository"
	"github.com/xuri/excelize/v2"
)

type ItemsFileService struct {
	categoryService CategoryService
	dataManager     repository.DataManager
	storeService    StoreService
}

func NewItemsFileService(
	categoryService CategoryService,
	dataManager repository.DataManager,
	storeService StoreService,
) ItemsFileService {
	return ItemsFileService{
		categoryService: categoryService,
		dataManager:     dataManager,
		storeService:    storeService,
	}
}

func (s ItemsFileService) Upload(context context.Context, storeUUID string, r io.Reader) error {
	file, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	defer file.Close()

	rows, err := file.GetRows("Items")
	if err != nil {
		return err
	}

	store, err := s.storeService.FindByUUID(context, storeUUID)
	if err != nil {
		return err
	}

	categories, err := s.categoryService.FindByStoreUUID(context, storeUUID)
	if err != nil {
		return err
	}
	categoriesMap := s.getCategoriesMap(categories)

	items := make([]entity.Item, 0)

	cols := make(map[string]int)
	for i, row := range rows {
		if i == 0 {
			cols = s.getColumnsIndex(row)
			continue
		}

		price, err := strconv.ParseFloat(row[cols[constant.ItemsFileLabelPrice]], 64)
		if err != nil {
			return err
		}

		item := entity.Item{
			Name:        row[cols[constant.ItemsFileLabelName]],
			Description: row[cols[constant.ItemsFileLabelDescription]],
			UUID:        uuid.NewString(),
			Price:       price,
			Category:    categoriesMap[row[cols[constant.ItemsFileLabelCategory]]],
			Store:       store,
		}

		items = append(items, item)
	}

	tx, err := s.dataManager.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, item := range items {
		_, err = s.dataManager.Item().Create(context, item, tx)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (s ItemsFileService) getColumnsIndex(row []string) map[string]int {
	labels := make(map[string]int)

	for i, col := range row {
		labels[col] = i
	}
	return labels
}

func (s ItemsFileService) getCategoriesMap(categories []entity.Category) map[string]entity.Category {
	r := make(map[string]entity.Category)

	for _, category := range categories {
		r[category.Name] = category
	}
	return r
}
