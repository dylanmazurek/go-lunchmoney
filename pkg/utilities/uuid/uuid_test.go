package uuid_test

import (
	"errors"
	"testing"

	"github.com/dylanmazurek/go-lunchmoney/pkg/utilities/uuid"
)

func TestParsing(t *testing.T) {
	tests := []struct {
		name  string
		input string

		wantUUIDStr    string
		wantParsedType string
		wantError      error
	}{
		{
			name:  "pass composed id to uuid",
			input: "lunch-portfolio-test-b1365644-f8c4-4241-957d-a4abfb87b2e8",

			wantUUIDStr:    "c2bac134-4002-5100-62b6-dd162d686eaf",
			wantParsedType: uuid.ParseTypeComposed,
		},
		{
			name:  "pass uuid to uuid",
			input: "d09b1ebb-7ed6-4f1c-b7a1-50138688a644",

			wantUUIDStr:    "d09b1ebb-7ed6-4f1c-b7a1-50138688a644",
			wantParsedType: uuid.ParseTypeUUID,
		},
		{
			name:  "pass raw uuid to uuid",
			input: "b1365644-f8cc-4241-957d-a4abfb87b2e8",

			wantUUIDStr:    "b1365644-f8cc-4241-957d-a4abfb87b2e8",
			wantParsedType: uuid.ParseTypeUUID,
		},
		{
			name:  "error empty uuid string",
			input: "",

			wantError: uuid.ErrEmptyString,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			uuid, errParse := uuid.Parse(test.input)

			uuidString, errString := uuid.String()

			err := errors.Join(errParse, errString)
			if err != nil {
				if test.wantError == nil {
					t.Errorf("test %q: got err %q, want nil", test.input, err)
				} else if !errors.Is(err, test.wantError) {
					t.Errorf("test %q: got %q, want err %q", test.input, err, test.wantError)
				}
			}

			if uuidString != nil && *uuidString != test.wantUUIDStr {
				t.Errorf("test %q: got %q, want %q", test.input, *uuidString, test.wantUUIDStr)
			}

			if uuid != nil && uuid.ParsedUsing != test.wantParsedType {
				t.Errorf("test %q: got %q, want %q", test.input, uuid.ParsedUsing, test.wantParsedType)
			}
		})
	}
}
