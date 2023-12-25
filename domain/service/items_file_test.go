package service

import (
	"testing"

	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestGetColumnsIndex(t *testing.T) {
	t.Parallel()

	service := ItemsFileService{}

	row := []string{"Name", "Description", "Price", "Category"}
	cols := service.getColumnsIndex(row)

	expected := map[string]int{
		"Name":        0,
		"Description": 1,
		"Price":       2,
		"Category":    3,
	}

	assert.EqualValues(t, expected, cols)
}

func TestGetCategoriesMap(t *testing.T) {
	t.Parallel()

	service := ItemsFileService{}

	category01 := entity.Category{Name: "Category 01"}
	category02 := entity.Category{Name: "Category 02"}

	categories := []entity.Category{
		category01,
		category02,
	}

	expected := map[string]entity.Category{
		category01.Name: category01,
		category02.Name: category02,
	}

	output := service.getCategoriesMap(categories)

	assert.EqualValues(t, expected, output)
}
