package transformation

import "testing"

func TestSingleWordQuote(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	got := QuoteReplace(input)
	expected := "I am exactly how they describe me: 'awesome'"

	if got != expected {
		t.Errorf("QuoteReplace(%q) = %q, expected = %q", input, got, expected)
	}
}

func TestMultipleWordQuote(t *testing.T) {
	input := "As Elton John said: ' I am the most well-known homosexual in the world '"
	got := QuoteReplace(input)
	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"

	if got != expected {
		t.Errorf("QuoteReplace(%q) = %q, expected = %q", input, got, expected)
	}
}
