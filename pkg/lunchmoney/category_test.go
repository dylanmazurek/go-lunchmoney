package lunchmoney

import (
	"net/http"
	"testing"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/stretchr/testify/assert"
)

func TestListCategory(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	expectedCategories := []models.Category{
		{ID: 1, Name: "Category 1"},
		{ID: 2, Name: "Category 2"},
	}

	// Mock the Do function
	client.Do = func(req *models.Request, resp interface{}) error {
		categoriesResponse := &models.CategoryResponse{
			Categories: expectedCategories,
		}
		*resp.(*models.CategoryResponse) = *categoriesResponse
		return nil
	}

	categories, err := client.ListCategory()
	assert.NoError(t, err)
	assert.Equal(t, &expectedCategories, categories)
}
