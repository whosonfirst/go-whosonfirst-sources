package sources

import (
	"testing"
)

func TestSpec(t *testing.T) {

	_, err := Spec()

	if err != nil {
		t.Fatalf("Failed to load spec, %v", err)
	}
}

func TestIsValidSource(t *testing.T) {

	valid := []string{
		"m49",
		"wof",
		"4sq",
		"arg-caba",
	}

	not_valid := []string{
		"omg",
		"wtf",
		"bbq",
	}

	for _, prefix := range valid {

		if !IsValidSource(prefix) {
			t.Fatalf("Expected '%s' to be valid", prefix)
		}
	}

	for _, prefix := range not_valid {

		if IsValidSource(prefix) {
			t.Fatalf("Did not expect '%s' to be valid", prefix)
		}
	}
}

func TestGetSourceByName(t *testing.T) {

	tests := map[string]int64{
		"arg-caba": 1108969549,
		"fb":       840464287,
		"4sq":      857075439,
	}

	for prefix, id := range tests {

		src, err := GetSourceByName(prefix)

		if err != nil {
			t.Fatalf("Unable to retrieve source by name '%s', %v", prefix, err)
		}

		if src.Id != id {
			t.Fatalf("Unexpected ID for '%s', expected '%d' but got '%d'", prefix, id, src.Id)
		}
	}

}
