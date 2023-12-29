//go:build unit
// +build unit

package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEqual(t *testing.T) {
	t.Parallel()

	item1 := Item{
		Code:        "ABC",
		Name:        "Item 1",
		Description: "Description 1",
		Price:       10.0,
		Category: Category{
			Name: "Category 1",
		},
	}
	item2 := Item{
		Code:        "ABC",
		Name:        "Item 1",
		Description: "Description 1",
		Price:       10.0,
		Category: Category{
			Name: "Category 1",
		},
	}
	item3 := Item{
		Code:        "XYZ",
		Name:        "Item 2",
		Description: "Description 2",
		Price:       20.0,
		Category: Category{
			Name: "Category 2",
		},
	}

	assert.True(t, item1.IsEqual(item2))
	assert.False(t, item1.IsEqual(item3))
}
