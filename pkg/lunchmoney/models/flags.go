package models

import "time"

type Flags struct {
	LunchmoneyAPIKey string
	AssetID          int64
	StartDate        time.Time
	EndDate          time.Time
}
