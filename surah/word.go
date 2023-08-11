// surah returns a surah along with it's word by word translation
package surah

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sort"

	"gihub.com/wizsk/quran-server/shared"
)

var ArabicDir string = "static/json/arabic/"

var (
	ErrLangNotFound = errors.New("surah: language not found")
)

func Surah(lang, idx string) (SurahByWord, error) {
	var err error
	if err = validSurah(lang, idx); err != nil {
		return SurahByWord{}, err
	}

	file, err := os.ReadFile(filepath.Join(ArabicDir, lang, idx+".json"))
	if err != nil {
		return SurahByWord{}, err
	}

	var surah SurahByWord
	err = json.Unmarshal(file, &surah)
	return surah, err
}

// returns availabe word by word langs
func AvailableLangs() ([]string, error) {
	dirs, err := os.ReadDir(ArabicDir)
	if err != nil {
		return nil, err
	}

	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name() < dirs[j].Name()
	})

	var langs []string
	for _, d := range dirs {
		if d.IsDir() {
			langs = append(langs, d.Name())
		}
	}
	return langs, nil
}

func validSurah(lang, idx string) error {
	langs, err := AvailableLangs()
	if err != nil {
		return err
	}

	var found bool
	for _, l := range langs {
		if lang == l {
			found = true
			break
		}
	}

	if !found {
		return ErrLangNotFound
	}

	return shared.IsValidIndex(idx)
}
