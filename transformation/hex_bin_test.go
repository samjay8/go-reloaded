package transformation

import "testing"

func TestHexConversion(t *testing.T) {
	input := "1E (hex) files were added"
	got := HexConversion(input)
	expected := "30 files were added"

	if got != expected {
		t.Errorf("HexConversion(%q) = %q, expected %q", input, got, expected)
	}
}

func TestBinConversion(t *testing.T) {
	input := "It has been 10 (bin) years"
	got := BinConversion(input)
	expected := "It has been 2 years"

	if got != expected {
		t.Errorf("BinConversion(%q) = %q, expected: %q", input, got, expected)
	}
}
