package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
	arabicFile, err := os.Open(StaticDir + "/json/arabic/" + index + ".json")
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
	surahInfoFile, err := os.Open(StaticDir + "/json/chapters/" + index + ".json")
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
	surahBangla, err := os.Open(StaticDir + "/json/bangla/" + index + ".json")
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
		translationFile, err := os.Open(StaticDir + "/json/" + path + index + ".json")
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

func GetWordByWord(index, lang string) ([]WordByWord, error) {
	var words []WordByWord
	path := StaticDir + "/json/word_by_word/" + lang + "/" + index + "/"
	pages, err := GetPageNumbers(path)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= pages; i++ {
		f := fmt.Sprintf("%s%d.json", path, i)
		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}

		var word WordByWord
		if err = json.NewDecoder(file).Decode(&word); err != nil {
			return nil, err
		}
		words = append(words, word)
		file.Close()
	}

	offset := words[0].Verses[0].Id - 1
	for i := range words {
		for j := range words[i].Verses {
			words[i].Verses[j].Id -= offset
		}

	}

	for i := range words {
		for j := range words[i].Verses {
			words[i].Verses[j].IndexForTrans = words[i].Verses[j].Id - 1
		}

	}

	return words, nil
}

// used fot he word by word cause it's borken into many pages..
func GetPageNumbers(path string) (int, error) {
	f, err := os.ReadFile(path + "page_count.txt")
	if err != nil {
		return -1, err
	}
	id := string(f)
	if id[len(id)-1] == '\n' {
		id = id[:len(id)-1]
	}
	pages, err := strconv.Atoi(string(id))
	if err != nil {
		return -1, err
	}
	return pages, nil
}
