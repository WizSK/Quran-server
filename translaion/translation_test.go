package translaion

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func init() {
	TranslaionDir = filepath.Join("..", TranslaionDir)
}

func TestRresourceId(t *testing.T) {
	translist, err := TranslationsList()
	if err != nil {
		t.Error(err)
		return
	}

	transNames := []string{}
	for _, tr := range translist {
		transNames = append(transNames, tr.TransLations...)
	}

	pat, err := translaionFilepaths(transNames...)
	if err != nil {
		t.Error(err)
		return
	}

	for _, p := range pat {
		f, err := os.ReadFile(filepath.Join(p, "1.json"))
		if err != nil {
			t.Error(err)
			return
		}
		var tran Translation
		err = json.Unmarshal(f, &tran)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(tran.Meta.Author, tran.Meta.Filters.ResourceId)
	}

}

func TestListTranslaions(t *testing.T) {
	_, err := TranslationsList()
	if err != nil {
		t.Error(err)
		return
	}
	// t.Logf("%#v\n", AvailableTranslation)
}

func TestTranslaionsPath(t *testing.T) {
	_, err := translaionFilepaths("Saheeh International")
	if err != nil {
		t.Error(err)
		return
	}
	// t.Logf("%#v\n", p)
}
