package translaion

import (
	"path/filepath"
	"testing"
)

func init() {
	TranslaionDir = filepath.Join("..", TranslaionDir)
}

func TestListTranslaions(t *testing.T) {
	AvailableTranslation, err := TranslationsList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v\n", AvailableTranslation)
}

func TestTranslaionsPath(t *testing.T) {
	p, err := translaionFilepaths("Saheeh International")
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%#v\n", p)
}
