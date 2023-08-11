package surah

import (
	"errors"
	"path/filepath"
	"testing"
)

// test will directly use the fileSystem. As the server is relient on that.
func init() {
	ArabicDir = filepath.Join("..", ArabicDir)
}

func TestAvailabeDirs(t *testing.T) {
	langs, err := AvailableLangs()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(langs)
}

func TestSurah(t *testing.T) {
	idx := "2"
	langs, err := AvailableLangs()
	if err != nil {
		t.Error(err)
	}

	for _, lang := range langs {
		_, err := Surah(lang, idx)
		if err != nil {
			t.Error(err)
		}
	}

}

func TestUnableToReadSurah(t *testing.T) {
	tests := []struct {
		lang string
		err  error
	}{
		// test index
		{
			lang: "foo",
			err:  ErrLangNotFound,
		},
		{
			lang: "bar",
			err:  ErrLangNotFound,
		},
		{
			lang: "",
			err:  ErrLangNotFound,
		},
	}

	idx := "1"
	for _, test := range tests {
		err := validSurah(test.lang, idx)
		if !errors.Is(err, test.err) {
			t.Errorf("expted %q; got %q", test.err, err)
		}
	}
}
