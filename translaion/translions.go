package translaion

import (
	"os"
	"path/filepath"
)

var TranslaionDir string = "static/json/translations/"

func GetTransLations(langs ...string) TransLations {
	return nil
}

func TranslationsList() (AvailableTranslations, error) {
	langs, err := os.ReadDir(TranslaionDir)
	if err != nil {
		return nil, err
	}

	var trans AvailableTranslations
	for _, lang := range langs {
		if !lang.IsDir() {
			continue
		}

		var tran AvailableTranslation
		tran.Language = lang.Name()

		// translations
		trns, err := os.ReadDir(filepath.Join(TranslaionDir, lang.Name()))
		if err != nil {
			return nil, err
		}
		for _, t := range trns {
			if !t.IsDir() {
				continue
			}
			tran.TransLations = append(tran.TransLations, t.Name())
		}
		trans = append(trans, tran)
	}

	return trans, nil
}

func translaionFilepaths(trans ...string) ([]string, error) {
	transList, err := TranslationsList()
	if err != nil {
		return nil, err
	}

	var path []string
	for _, tran := range trans {
		for _, lang := range transList {
			for _, transName := range lang.TransLations {
				if tran == transName {
					path = append(path, filepath.Join(TranslaionDir, lang.Language, transName))
				}
			}
		}
	}

	return path, nil
}
