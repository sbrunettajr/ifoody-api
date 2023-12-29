package service

import (
	"bytes"
	"context"
	"errors"
	"fmt"
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
	itemService     ItemService
	storeService    StoreService
}

func NewItemsFileService(
	categoryService CategoryService,
	dataManager repository.DataManager,
	itemService ItemService,
	storeService StoreService,
) ItemsFileService {
	return ItemsFileService{
		categoryService: categoryService,
		dataManager:     dataManager,
		itemService:     itemService,
		storeService:    storeService,
	}
}

// Add integration test
func (s ItemsFileService) Upload(context context.Context, storeUUID string, r io.Reader) error {
	file, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}
	defer file.Close()

	rows, err := file.GetRows(constant.ItemsFileItemsSheetName)
	if err != nil {
		return err
	}

	items, err := s.dataManager.Item().FindByStoreUUIDWithRelations(context, storeUUID)
	if err != nil {
		return err
	}
	itemsMap := s.getItemsMapByCode(items)

	categories, err := s.categoryService.FindByStoreUUID(context, storeUUID)
	if err != nil {
		return err
	}
	categoriesMap := s.getCategoriesMapByUUID(categories)

	store, err := s.storeService.FindByUUID(context, storeUUID)
	if err != nil {
		return err
	}

	labels := make(map[string]int)
	insertItems := make([]entity.Item, 0)

	for i, row := range rows {
		if i == 0 {
			labels = s.getLabelsIndex(row)

			err = s.validateLabels(labels)
			if err != nil {
				return err
			}
			continue
		}

		price, err := strconv.ParseFloat(row[labels[constant.ItemsFileLabelPrice]], 64)
		if err != nil {
			return err
		}

		category, ok := categoriesMap[row[labels[constant.ItemsFileLabelCategory]]]
		if !ok {
			return errors.New("") // Add error message
		}

		item := entity.Item{
			Code:        row[labels[constant.ItemsFileLabelCode]],
			Name:        row[labels[constant.ItemsFileLabelName]],
			Description: row[labels[constant.ItemsFileLabelDescription]],
			UUID:        uuid.NewString(),
			Price:       price,
			Category:    category,
			Store:       store,
		}

		if i, ok := itemsMap[item.Code]; ok {
			if item.IsEqual(i) {
				continue
			}

			item.ID = i.ID

			err = s.itemService.Update(context, item)
			if err != nil {
				return err
			}
			continue
		}
		insertItems = append(insertItems, item)
	}

	if len(insertItems) > 0 {
		err = s.dataManager.Item().BulkInsert(context, insertItems)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s ItemsFileService) getLabelsIndex(row []string) map[string]int {
	labels := make(map[string]int)

	for i, col := range row {
		labels[col] = i
	}
	return labels
}

func (s ItemsFileService) validateLabels(labelsMap map[string]int) error {
	labels := []string{
		constant.ItemsFileLabelCode,
		constant.ItemsFileLabelName,
		constant.ItemsFileLabelDescription,
		constant.ItemsFileLabelPrice,
		constant.ItemsFileLabelCategory,
	}

	if len(labelsMap) != len(labels) {
		return errors.New("") // Add error message
	}

	for _, label := range labels {
		if _, ok := labelsMap[label]; !ok {
			return errors.New("") // Add error message
		}
	}
	return nil
}

func (s ItemsFileService) getCategoriesMapByUUID(categories []entity.Category) map[string]entity.Category {
	r := make(map[string]entity.Category)

	for _, category := range categories {
		r[category.Name] = category
	}
	return r
}

func (s ItemsFileService) getItemsMapByCode(items []entity.Item) map[string]entity.Item {
	i := make(map[string]entity.Item)

	for _, item := range items {
		i[item.Code] = item
	}
	return i
}

// Add unit test
func (s ItemsFileService) setLabels(file *excelize.File) {
	style, _ := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	file.SetCellValue(constant.ItemsFileItemsSheetName, "A1", constant.ItemsFileLabelCode)
	file.SetCellValue(constant.ItemsFileItemsSheetName, "B1", constant.ItemsFileLabelName)
	file.SetCellValue(constant.ItemsFileItemsSheetName, "C1", constant.ItemsFileLabelDescription)
	file.SetCellValue(constant.ItemsFileItemsSheetName, "D1", constant.ItemsFileLabelPrice)
	file.SetCellValue(constant.ItemsFileItemsSheetName, "E1", constant.ItemsFileLabelCategory)

	file.SetCellStyle(constant.ItemsFileItemsSheetName, "A1", "E1", style)
}

// Add unit test
func (s ItemsFileService) makeCategoriesSheet(context context.Context, file *excelize.File, categories []entity.Category) error {
	index, err := file.NewSheet(constant.ItemsFileCategoriesSheetName)
	if err != nil {
		return err
	}

	file.SetActiveSheet(index)

	for i, category := range categories {
		name, err := excelize.CoordinatesToCellName(1, i+1)
		if err != nil {
			return err
		}

		file.SetCellValue(constant.ItemsFileCategoriesSheetName, name, category.Name)
	}
	return nil
}

// Add unit test
func (s ItemsFileService) makeDropList(file *excelize.File, categories []entity.Category) {
	dataValidation := excelize.NewDataValidation(true)
	dataValidation.Sqref = "E2:E1000"

	dataValidation.SetSqrefDropList(fmt.Sprintf("%s!$A$1:$A$%d", constant.ItemsFileCategoriesSheetName, len(categories)))
	file.AddDataValidation(constant.ItemsFileItemsSheetName, dataValidation)
}

// Add unit test
func (s ItemsFileService) fillItemsSheet(file *excelize.File, items []entity.Item) {
	for i, item := range items {
		file.SetCellValue(constant.ItemsFileItemsSheetName, fmt.Sprintf("A%d", i), item.Code)
		file.SetCellValue(constant.ItemsFileItemsSheetName, fmt.Sprintf("B%d", i), item.Name)
		file.SetCellValue(constant.ItemsFileItemsSheetName, fmt.Sprintf("C%d", i), item.Description)
		file.SetCellValue(constant.ItemsFileItemsSheetName, fmt.Sprintf("D%d", i), item.Price)
		file.SetCellValue(constant.ItemsFileItemsSheetName, fmt.Sprintf("E%d", i), item.Category.Name)
	}
}

// Add integration test
func (s ItemsFileService) Download(context context.Context, storeUUID string, isTemplate bool) ([]byte, error) {
	file := excelize.NewFile()
	defer file.Close()

	s.setLabels(file)

	categories, err := s.categoryService.FindByStoreUUID(context, storeUUID)
	if err != nil {
		return nil, err
	}

	err = s.makeCategoriesSheet(context, file, categories)
	if err != nil {
		return nil, err
	}

	file.SetActiveSheet(constant.ItemsFileItemsSheetIndex)

	err = file.SetSheetVisible(constant.ItemsFileCategoriesSheetName, false)
	if err != nil {
		return nil, err
	}

	err = file.SetSheetName(constant.DefaultSheetName, constant.ItemsFileItemsSheetName)
	if err != nil {
		return nil, err
	}

	s.makeDropList(file, categories)

	if !isTemplate {
		items, err := s.itemService.FindByStoreUUIDWithRelations(context, storeUUID)
		if err != nil {
			return nil, err
		}

		s.fillItemsSheet(file, items)
	}

	var buffer bytes.Buffer
	file.Write(&buffer)

	return buffer.Bytes(), nil
}
