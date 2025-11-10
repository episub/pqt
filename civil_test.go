package pqt

import (
	"encoding/json"
	"testing"

	"cloud.google.com/go/civil"
)

func TestAddMonths(t *testing.T) {
	type Case struct {
		Date           Date
		Add            int
		ExpectedString string
	}
	cases := []Case{
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            5,
			ExpectedString: "2000-06-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            36,
			ExpectedString: "2003-01-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            14,
			ExpectedString: "2001-03-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            -1,
			ExpectedString: "1999-12-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            -12,
			ExpectedString: "1999-01-01",
		},
		{
			Date:           Date{Date: civil.Date{Day: 1, Month: 1, Year: 2000}},
			Add:            -14,
			ExpectedString: "1998-11-01",
		},
	}

	for _, c := range cases {
		c.Date.AddMonths(c.Add)
		if c.Date.String() != c.ExpectedString {
			t.Errorf("Expected string %s, but had %s", c.ExpectedString, c.Date.String())
		}
	}
}

func TestNullDate_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		nullDate NullDate
		expected string
	}{
		{
			name: "valid date",
			nullDate: NullDate{
				Date:  civil.Date{Year: 2023, Month: 12, Day: 25},
				Valid: true,
			},
			expected: `"2023-12-25"`,
		},
		{
			name: "null date",
			nullDate: NullDate{
				Valid: false,
			},
			expected: "null",
		},
		{
			name: "zero date but valid",
			nullDate: NullDate{
				Date:  civil.Date{},
				Valid: true,
			},
			expected: `"0000-00-00"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := json.Marshal(tt.nullDate)
			if err != nil {
				t.Errorf("MarshalJSON() error = %v", err)
				return
			}
			if string(result) != tt.expected {
				t.Errorf("MarshalJSON() = %v, want %v", string(result), tt.expected)
			}
		})
	}
}

func TestNullDate_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected NullDate
		wantErr  bool
	}{
		{
			name:  "valid date string",
			input: `"2023-12-25"`,
			expected: NullDate{
				Date:  civil.Date{Year: 2023, Month: 12, Day: 25},
				Valid: true,
			},
			wantErr: false,
		},
		{
			name:  "null value",
			input: "null",
			expected: NullDate{
				Valid: false,
			},
			wantErr: false,
		},
		{
			name:  "empty string",
			input: `""`,
			expected: NullDate{
				Valid: false,
			},
			wantErr: false,
		},
		{
			name:    "nil data",
			input:   "",
			wantErr: true,
		},
		{
			name:    "invalid date format",
			input:   `"invalid-date"`,
			wantErr: true,
		},
		{
			name:    "invalid JSON",
			input:   `invalid-json`,
			wantErr: true,
		},
		{
			name:  "leap year date",
			input: `"2024-02-29"`,
			expected: NullDate{
				Date:  civil.Date{Year: 2024, Month: 2, Day: 29},
				Valid: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var nd NullDate
			err := json.Unmarshal([]byte(tt.input), &nd)

			if tt.wantErr {
				if err == nil {
					t.Errorf("UnmarshalJSON() expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("UnmarshalJSON() unexpected error = %v", err)
				return
			}

			if nd.Valid != tt.expected.Valid {
				t.Errorf("UnmarshalJSON() Valid = %v, want %v", nd.Valid, tt.expected.Valid)
			}

			if nd.Valid && nd.Date != tt.expected.Date {
				t.Errorf("UnmarshalJSON() Date = %v, want %v", nd.Date, tt.expected.Date)
			}
		})
	}
}

func TestNullDate_MarshalUnmarshalRoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		original NullDate
	}{
		{
			name: "valid date roundtrip",
			original: NullDate{
				Date:  civil.Date{Year: 2023, Month: 6, Day: 15},
				Valid: true,
			},
		},
		{
			name: "invalid date roundtrip",
			original: NullDate{
				Valid: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Marshal
			data, err := json.Marshal(tt.original)
			if err != nil {
				t.Errorf("Marshal error = %v", err)
				return
			}

			// Unmarshal
			var unmarshaled NullDate
			err = json.Unmarshal(data, &unmarshaled)
			if err != nil {
				t.Errorf("Unmarshal error = %v", err)
				return
			}

			// Compare
			if unmarshaled.Valid != tt.original.Valid {
				t.Errorf("Roundtrip Valid = %v, want %v", unmarshaled.Valid, tt.original.Valid)
			}

			if unmarshaled.Valid && unmarshaled.Date != tt.original.Date {
				t.Errorf("Roundtrip Date = %v, want %v", unmarshaled.Date, tt.original.Date)
			}
		})
	}
}

func TestNullDate_InStructContext(t *testing.T) {
	type TestRecord struct {
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		BirthDate NullDate `json:"birth_date"`
	}

	tests := []struct {
		name     string
		input    string
		expected TestRecord
		wantErr  bool
	}{
		{
			name:  "struct with valid date",
			input: `{"id": 1, "name": "John", "birth_date": "1990-05-15"}`,
			expected: TestRecord{
				ID:   1,
				Name: "John",
				BirthDate: NullDate{
					Date:  civil.Date{Year: 1990, Month: 5, Day: 15},
					Valid: true,
				},
			},
			wantErr: false,
		},
		{
			name:  "struct with null date",
			input: `{"id": 2, "name": "Jane", "birth_date": null}`,
			expected: TestRecord{
				ID:   2,
				Name: "Jane",
				BirthDate: NullDate{
					Valid: false,
				},
			},
			wantErr: false,
		},
		{
			name:  "struct with missing date field",
			input: `{"id": 3, "name": "Bob"}`,
			expected: TestRecord{
				ID:   3,
				Name: "Bob",
				BirthDate: NullDate{
					Valid: false,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var record TestRecord
			err := json.Unmarshal([]byte(tt.input), &record)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error, but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error = %v", err)
				return
			}

			if record.ID != tt.expected.ID {
				t.Errorf("ID = %v, want %v", record.ID, tt.expected.ID)
			}

			if record.Name != tt.expected.Name {
				t.Errorf("Name = %v, want %v", record.Name, tt.expected.Name)
			}

			if record.BirthDate.Valid != tt.expected.BirthDate.Valid {
				t.Errorf("BirthDate.Valid = %v, want %v", record.BirthDate.Valid, tt.expected.BirthDate.Valid)
			}

			if record.BirthDate.Valid && record.BirthDate.Date != tt.expected.BirthDate.Date {
				t.Errorf("BirthDate.Date = %v, want %v", record.BirthDate.Date, tt.expected.BirthDate.Date)
			}

			// Test round-trip marshaling
			marshaled, err := json.Marshal(record)
			if err != nil {
				t.Errorf("Marshal error = %v", err)
				return
			}

			var roundTrip TestRecord
			err = json.Unmarshal(marshaled, &roundTrip)
			if err != nil {
				t.Errorf("Round-trip unmarshal error = %v", err)
				return
			}

			if roundTrip.BirthDate.Valid != record.BirthDate.Valid {
				t.Errorf("Round-trip BirthDate.Valid = %v, want %v", roundTrip.BirthDate.Valid, record.BirthDate.Valid)
			}

			if roundTrip.BirthDate.Valid && roundTrip.BirthDate.Date != record.BirthDate.Date {
				t.Errorf("Round-trip BirthDate.Date = %v, want %v", roundTrip.BirthDate.Date, record.BirthDate.Date)
			}
		})
	}
}
