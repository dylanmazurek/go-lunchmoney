package date_test

import (
	"testing"
	"time"

	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/date"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		input string

		wantDate date.Date
		wantJson *string
		wantErr  error
	}{
		{
			name:  "success date only",
			input: "2023-01-01",

			wantDate: date.Date{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)},
		},
		{
			name:  "success date only 2",
			input: "2023-01-02",

			wantDate: date.Date{
				Date: time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name:  "success json marshal",
			input: "2023-01-03",

			wantDate: date.Date{
				Date: time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC),
			},
			wantJson: &[]string{"{\"Date\":\"2023-01-03T00:00:00Z\"}"}[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedDate, err := date.Parse(tt.input)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			if parsedDate != tt.wantDate {
				t.Errorf("Parse() = %v, want %v", parsedDate, tt.wantDate)
			}

			if tt.wantJson != nil {
				jsonBytes, err := parsedDate.MarshalJSON()
				if err != nil {
					t.Errorf("Parse() error = %v", err)
				}

				jsonStr := string(jsonBytes)
				if jsonStr != *tt.wantJson {
					t.Errorf("MarshalJSON() = %s, want %s", parsedDate, *tt.wantJson)
				}
			}

		})
	}
}

func TestParseDate(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time

		wantDate date.Date
		wantJson *string
		wantErr  error
	}{
		{
			name: "success date time",
			input: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2023-01-04T01:22:00Z")
				return t
			}(),

			wantDate: date.Date{time.Date(2023, 1, 4, 0, 0, 0, 0, time.UTC)},
		},
		{
			name: "success json marshal",
			input: func() time.Time {
				t, _ := time.Parse(time.RFC3339, "2025-07-04T04:28:00Z")
				return t
			}(),

			wantDate: date.Date{
				Date: time.Date(2025, 7, 4, 0, 0, 0, 0, time.UTC),
			},
			wantJson: &[]string{"{\"Date\":\"2025-07-04T00:00:00Z\"}"}[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedDate, err := date.ParseDate(tt.input)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("ParseDate() error = %v, wantErr %v", err, tt.wantErr)
				}

				return
			}

			if parsedDate != tt.wantDate {
				t.Errorf("ParseDate() = %v, want %v", parsedDate, tt.wantDate)
			}

			if tt.wantJson != nil {
				jsonBytes, err := parsedDate.MarshalJSON()
				if err != nil {
					t.Errorf("ParseDate() error = %v", err)
				}

				jsonStr := string(jsonBytes)
				if jsonStr != *tt.wantJson {
					t.Errorf("MarshalJSON() = %s, want %s", parsedDate, *tt.wantJson)
				}
			}

		})
	}
}
