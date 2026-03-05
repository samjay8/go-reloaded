package transformation

import "testing"

func TestPunctuation(t *testing.T) {
	input := "hello , world"
	got := Punctuation(input)
	expected := "hello, world"

	if got != expected {
		t.Errorf("Punctuation(%q) = %q, expected = %q", input, got, expected)
	}
}

func TestAfterPunctuation(t *testing.T) {
	input := "hello,world"
	got := Punctuation(input)
	expected := "hello, world"

	if got != expected {
		t.Errorf("Punctuation(%q) = %q, expected = %q", input, got, expected)
	}

}

func TestGroupedPunctuation(t *testing.T) {
	input := "I was thinking ... You were right"
	got := Punctuation(input)
	expected := "I was thinking... You were right"

	if got != expected {
		t.Errorf("Punctuation(%q) = %q, expected = %q", input, got, expected)
	}
}

func TestMultiplePunctuation(t *testing.T) {
	input := "hello , world ! how are you ?"
	got := Punctuation(input)
	expected := "hello, world! how are you?"

	if got != expected {
		t.Errorf("Punctuation(%q) = %q, expected = %q", input, got, expected)
	}
}
