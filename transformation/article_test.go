package transformation

import "testing"

func TestArticleReplace(t *testing.T) {
	input := "There it was. A amazing rock!"
	got := ArticleReplace(input)
	expected := "There it was. An amazing rock!"

	if got != expected {
		t.Errorf("ArticleReplace(%q) = %q, expected %q", input, got, expected)
	}
}
