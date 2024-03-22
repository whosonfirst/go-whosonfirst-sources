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
		// by prefix
		"m49",
		"wof",
		"4sq",
		"arg-caba",
		// by name
		"metazen",
	}

	not_valid := []string{
		"omg",
		"wtf",
		"bbq",
	}

	for _, label := range valid {

		if !IsValidSource(label) {
			t.Fatalf("Expected '%s' to be valid", label)
		}
	}

	for _, label := range not_valid {

		if IsValidSource(label) {
			t.Fatalf("Did not expect '%s' to be valid", label)
		}
	}
}

func TestGetSourceByPrefix(t *testing.T) {

	tests := map[string]int64{
		"arg-caba": 1108969549,
		"fb":       840464287,
		"4sq":      857075439,
	}

	for prefix, id := range tests {

		src, err := GetSourceByPrefix(prefix)

		if err != nil {
			t.Fatalf("Unable to retrieve source by prefix '%s', %v", prefix, err)
		}

		if src.Id != id {
			t.Fatalf("Unexpected ID for '%s', expected '%d' but got '%d'", prefix, id, src.Id)
		}
	}

}

func TestGetSourceByName(t *testing.T) {

	tests := map[string]int64{
		"arg-caba":   1108969549,
		"metazen":    404734197,
		"foursquare": 857075439,
	}

	for name, id := range tests {

		src, err := GetSourceByName(name)

		if err != nil {
			t.Fatalf("Unable to retrieve source by name '%s', %v", name, err)
		}

		if src.Id != id {
			t.Fatalf("Unexpected ID for '%s', expected '%d' but got '%d'", name, id, src.Id)
		}
	}

}
