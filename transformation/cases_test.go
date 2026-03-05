package transformation

import "testing"

func TestUpperCase(t *testing.T) {
	input := "Ready, set, go (up) !"
	got := UpperCase(input)
	expected := "Ready, set, GO !"

	if got != expected {
		t.Errorf("UpperCase(%q) = %q, expected %q", input, got, expected)
	}
}

func TestLowerCase(t *testing.T) {
	input := "I should stop SHOUTING (low)"
	got := LowerCase(input)
	expected := "I should stop shouting"

	if got != expected {
		t.Errorf("LowerCase(%q) = %q, expected %q", input, got, expected)
	}
}

func TestCapCase(t *testing.T) {
	input := "it (cap) was the best of times"
	got := CapCase(input)
	expected := "It was the best of times"
	if got != expected {
		t.Errorf("CapCase(%q) = %q, expected %q", input, got, expected)
	}
}
