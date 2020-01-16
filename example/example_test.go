package example

import (
	"database/sql/driver"
	"strings"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		check    Enum1
		expected bool
	}{
		{check: Enum1_PROFILE, expected: true},
		{check: Enum1_BILLING, expected: true},
		{check: Enum1_LEGAL, expected: true},
		{check: 4, expected: false},
	}
	for _, test := range tests {
		if test.check.IsValid() != test.expected {
			t.Errorf("Unexpected IsValid for enum value %d", test.check)
		}
	}
}

func TestDatabaseValue(t *testing.T) {
	tests := []struct {
		check    Enum1
		expected driver.Value
		hasErr   bool
	}{
		{Enum1_PROFILE, "snake", false},
		{Enum1_BILLING, "BILLING", false},
		{Enum1_LEGAL, "LEGAL", false},
		{4, nil, true},
	}
	for _, test := range tests {
		v, err := test.check.Value()
		if (err != nil) != test.hasErr {
			t.Errorf("Unexpected error from Enum1.Value: %s", err.Error())
		}
		if v != test.expected {
			t.Errorf("Unexpected database value for enum value %d got %v", test.check, v)
		}
	}
}

func TestDatabaseScan(t *testing.T) {
	tests := []struct {
		check    string
		expected Enum1
		hasErr   bool
	}{
		{"snake", Enum1_PROFILE, false},
		{"BILLING", Enum1_BILLING, false},
		{"LEGAL", Enum1_LEGAL, false},
		{"lol", Enum1_PROFILE, true}, // Zeroth enum values are default :\
	}
	for _, test := range tests {
		var e Enum1
		err := e.Scan(test.check)
		if (err != nil) != test.hasErr {
			t.Errorf("Unexpected error from Enum1.Value: %s", err.Error())
		}
		if e != test.expected {
			t.Errorf("Unexpected database scan for value %s got %v", test.check, e)
		}
	}
}

func TestMarshalGQL(t *testing.T) {
	tests := []struct {
		check    Enum1
		expected string
		hasErr   bool
	}{
		{Enum1_PROFILE, `"turkey"`, false},
		{Enum1_BILLING, `"BILLING"`, false},
		{Enum1_LEGAL, `"blarg"`, false},
		{4, `""`, true},
	}
	for _, test := range tests {
		b := &strings.Builder{}
		test.check.MarshalGQL(b)
		if b.String() != test.expected {
			t.Errorf("Unexpected graphql value for enum value %d got %s", test.check, b.String())
		}
	}
}

func TestUnmarshalGQL(t *testing.T) {
	tests := []struct {
		check    string
		expected Enum1
		hasErr   bool
	}{
		{"turkey", Enum1_PROFILE, false},
		{"BILLING", Enum1_BILLING, false},
		{"blarg", Enum1_LEGAL, false},
		{"lol", Enum1_PROFILE, true}, // Zeroth enum values are default :\
	}
	for _, test := range tests {
		var e Enum1
		err := e.UnmarshalGQL(test.check)
		if (err != nil) != test.hasErr {
			t.Errorf("Unexpected error from Enum1.Value: %s", err.Error())
		}
		if e != test.expected {
			t.Errorf("Unexpected UnmarshalGQL for value %s got %v", test.check, e)
		}
	}
}
