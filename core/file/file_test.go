package file

import "testing"

func TestLanguageFileExtension(t *testing.T) {
	if got := LanguageFileExtension("python3"); got != ".py" {
		t.Errorf("python3 -> %q", got)
	}
	if got := LanguageFileExtension("go"); got != ".go" {
		t.Errorf("go -> %q (want .go via golang alias in config)", got)
	}
	if got := LanguageFileExtension("unknown-lang"); got != ".txt" {
		t.Errorf("unknown -> %q", got)
	}
}
