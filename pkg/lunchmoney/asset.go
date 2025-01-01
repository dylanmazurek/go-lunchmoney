package lunchmoney

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/constants"
	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
)

// FetchAsset fetches an asset by its ID.
func (c *Client) FetchAsset(assetId int64) (*models.Asset, error) {
	assets, err := c.ListAsset()
	if err != nil {
		return nil, err
	}

	assetIdx := slices.IndexFunc(*assets, func(asset models.Asset) bool { return asset.AssetID == &assetId })
	if assetIdx == -1 {
		return nil, nil
	}

	asset := (*assets)[assetIdx]

	return &asset, err
}

// ListAsset lists all assets.
func (c *Client) ListAsset() (*[]models.Asset, error) {
	req, err := c.NewRequest(http.MethodGet, constants.API_PATH_ASSETS, nil, nil)
	if err != nil {
		return nil, err
	}

	var assets models.AssetResponse
	err = c.Do(req, &assets)

	return &assets.Assets, err
}

// UpdateAsset updates an asset by its ID.
func (c *Client) UpdateAsset(id int64, asset *models.Asset) (*models.Asset, error) {
	assetJson, err := json.Marshal(asset)
	if err != nil {
		return nil, err
	}

	requestPath := fmt.Sprintf("%s/%d", constants.API_PATH_ASSETS, id)

	req, err := c.NewRequest(http.MethodPut, requestPath, bytes.NewReader(assetJson), nil)
	if err != nil {
		return nil, err
	}

	var updatedAsset models.Asset
	err = c.Do(req, &updatedAsset)

	if err != nil {
		return nil, err
	}

	return &updatedAsset, err
}
