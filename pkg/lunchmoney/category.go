package lunchmoney

import (
	"net/http"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/constants"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
)

func (c *Client) ListCategory() (*[]models.Category, error) {
	req, err := c.NewRequest(http.MethodGet, constants.API_PATH_CATEGORIES, nil, nil)
	if err != nil {
		return nil, err
	}

	var categories models.CategoryResponse
	err = c.Do(req, &categories)

	return &categories.Categories, err
}
