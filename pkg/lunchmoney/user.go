package lunchmoney

import (
	"net/http"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/constants"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
)

func (c *Client) Me() (*models.User, error) {
	req, err := c.NewRequest(http.MethodGet, constants.API_PATH_ME, nil, nil)
	if err != nil {
		return nil, err
	}

	var me models.User
	err = c.Do(req, &me)

	return &me, err
}
