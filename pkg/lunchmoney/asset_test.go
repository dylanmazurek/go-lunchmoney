package lunchmoney

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dylanmazurek/go-lunchmoney/pkg/lunchmoney/models"
	"github.com/stretchr/testify/assert"
)

func TestFetchAsset(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	assetID := int64(1)
	expectedAsset := &models.Asset{
		AssetID: &assetID,
	}

	// Mock the ListAsset function
	client.ListAsset = func() (*[]models.Asset, error) {
		return &[]models.Asset{*expectedAsset}, nil
	}

	asset, err := client.FetchAsset(assetID)
	assert.NoError(t, err)
	assert.Equal(t, expectedAsset, asset)
}

func TestListAsset(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	expectedAssets := []models.Asset{
		{AssetID: int64Ptr(1)},
		{AssetID: int64Ptr(2)},
	}

	// Mock the Do function
	client.Do = func(req *models.Request, resp interface{}) error {
		assetsResponse := &models.AssetResponse{
			Assets: expectedAssets,
		}
		*resp.(*models.AssetResponse) = *assetsResponse
		return nil
	}

	assets, err := client.ListAsset()
	assert.NoError(t, err)
	assert.Equal(t, &expectedAssets, assets)
}

func TestUpdateAsset(t *testing.T) {
	client := &Client{
		internalClient: &http.Client{},
	}

	assetID := int64(1)
	updatedAsset := &models.Asset{
		AssetID: &assetID,
	}

	// Mock the Do function
	client.Do = func(req *models.Request, resp interface{}) error {
		*resp.(*models.Asset) = *updatedAsset
		return nil
	}

	asset, err := client.UpdateAsset(assetID, updatedAsset)
	assert.NoError(t, err)
	assert.Equal(t, updatedAsset, asset)
}

func int64Ptr(i int64) *int64 {
	return &i
}
