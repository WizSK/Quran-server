package main

import (
	"encoding/json"
	"os"
)

// translations file structure
/*

static/json/
├─ english/clear_quran/
│  ├─ 1.json
│  ├─ 2.json
│  ├─ ...json
│
├─ Your Language
│  ├─ 1.json
│  ├─ 2.json
│  ├─ ...json

*/

// here static/json/
// var tnaslaitonList = []string{"english/clear_quran/", "english/saheeh_internatioanl/"}
var tnaslaitonList = []string{"english/clear_quran/"}

func GetArabicAyas(index string) (Ayas, error) {
	var ayas Ayas
	arabicFile, err := os.Open("static/json/arabic/" + index + ".json")
	if err != nil {
		return ayas, err
	}
	defer arabicFile.Close()

	if err = json.NewDecoder(arabicFile).Decode(&ayas); err != nil {
		return ayas, err
	}

	offset := ayas.Verses[0].Id - 1
	for i := 0; i < len(ayas.Verses); i++ {
		ayas.Verses[i].Id = ayas.Verses[i].Id - offset
	}

	return ayas, err
}

func GetSurahInfo(index string) (ChapterInfo, error) {
	var surahInfo ChapterInfo
	surahInfoFile, err := os.Open("static/json/chapters/" + index + ".json")
	if err != nil {
		return surahInfo, err
	}
	defer surahInfoFile.Close()

	// info
	if err = json.NewDecoder(surahInfoFile).Decode(&surahInfo); err != nil {
		return surahInfo, err
	}

	return surahInfo, err
}

func GetBanglaTranslation(index string) (TranslatedVerses, error) {
	var banglaTrans TranslatedVerses
	surahBangla, err := os.Open("static/json/bangla/" + index + ".json")
	if err != nil {
		return banglaTrans, err
	}
	defer surahBangla.Close()

	if err = json.NewDecoder(surahBangla).Decode(&banglaTrans); err != nil {
		return banglaTrans, err
	}

	return banglaTrans, nil
}

func GetTransLations(index string) ([]TranslatedVerses, error) {
	var translations []TranslatedVerses
	// Trnaslations
	for i, path := range tnaslaitonList {
		translationFile, err := os.Open("static/json/" + path + index + ".json")
		if err != nil {
			return nil, err
		}

		translations = append(translations, TranslatedVerses{})
		if err = json.NewDecoder(translationFile).Decode(&translations[i]); err != nil {
			translationFile.Close()
			return nil, err
		}
		translationFile.Close()
	}

	return translations, nil
}
