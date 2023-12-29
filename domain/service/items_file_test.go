//go:build unit
// +build unit

package service

import (
	"testing"

	"github.com/sbrunettajr/ifoody-api/domain/constant"
	"github.com/sbrunettajr/ifoody-api/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestGetLabelsIndex(t *testing.T) {
	t.Parallel()

	service := ItemsFileService{}

	row := []string{"Name", "Description", "Price", "Category"}
	cols := service.getLabelsIndex(row)

	expected := map[string]int{
		"Name":        0,
		"Description": 1,
		"Price":       2,
		"Category":    3,
	}

	assert.EqualValues(t, expected, cols)
}

func TestGetCategoriesMapByUUID(t *testing.T) {
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

	output := service.getCategoriesMapByUUID(categories)

	assert.EqualValues(t, expected, output)
}

func TestValidateLabels(t *testing.T) {
	t.Parallel()

	service := ItemsFileService{}

	testCases := []struct {
		desc      string
		labels    map[string]int
		expectErr bool
	}{
		{
			desc: "ShouldExecuteWithoutError",
			labels: map[string]int{
				constant.ItemsFileLabelCode:        1,
				constant.ItemsFileLabelName:        1,
				constant.ItemsFileLabelDescription: 1,
				constant.ItemsFileLabelPrice:       1,
				constant.ItemsFileLabelCategory:    1,
			},
			expectErr: false,
		},
		{
			desc: "ShouldExecuteWithError",
			labels: map[string]int{
				constant.ItemsFileLabelCode:        1,
				constant.ItemsFileLabelName:        1,
				constant.ItemsFileLabelDescription: 1,
				constant.ItemsFileLabelPrice:       1,
			},
			expectErr: true,
		},
	}

	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			err := service.validateLabels(tC.labels)
			assert.Equal(t, tC.expectErr, err != nil)
		})
	}
}

func TestGetItemsMapByCode(t *testing.T) {
	t.Parallel()

	service := ItemsFileService{}

	items := []entity.Item{
		{
			Code: "ABC",
			Name: "Item 1",
		},
		{
			Code: "DEF",
			Name: "Item 2",
		},
		{
			Code: "GHI",
			Name: "Item 3",
		},
	}
	itemsMap := service.getItemsMapByCode(items)

	assert.Len(t, itemsMap, 3)

	assert.Equal(t, itemsMap["ABC"].Name, "Item 1")
	assert.Equal(t, itemsMap["DEF"].Name, "Item 2")
	assert.Equal(t, itemsMap["GHI"].Name, "Item 3")
}
