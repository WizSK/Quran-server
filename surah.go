package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

// Cashing thing..
var SurahCash = make(map[string][]byte)

func getSurah(w http.ResponseWriter, r *http.Request, idx string) {
	id, err := strconv.Atoi(idx)

	if err != nil || id < 1 || id > 144 {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return
	}

	if _, ok := SurahCash[idx]; ok {
		w.Write(SurahCash[idx])
		return
	}

	surahArabic, err := os.Open("static/json/arabic/" + idx + ".json")
	if err != nil {
		w.Write([]byte("Page Not found"))
		fmt.Println(err)
		return
	}
	defer surahArabic.Close()

	surahInfo, err := os.Open("static/json/chapters/" + idx + ".json")

	if err != nil {
		w.Write([]byte("Page Not found"))
		fmt.Println(err)
		return
	}
	defer surahInfo.Close()

	var all CompleteSurah
	// Trnaslations
	for i, path := range tnaslaitonList {
		translation, err := os.Open("static/json/" + path + idx + ".json")
		if err != nil {
			w.Write([]byte("Page Not found"))
			fmt.Println(err)
			return
		}

		all.Translaions = append(all.Translaions, TranslatedVerses{})
		if err = json.NewDecoder(translation).Decode(&all.Translaions[i]); err != nil {
			fmt.Println(err)
			return
		}

		translation.Close()
	}

	surahBangla, err := os.Open("static/json/bangla/" + idx + ".json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer surahBangla.Close()

	// ara
	if err = json.NewDecoder(surahArabic).Decode(&all.Aarabic); err != nil {
		fmt.Println(err)
		return
	}

	// info
	if err = json.NewDecoder(surahInfo).Decode(&all.SurahInfo); err != nil {
		fmt.Println(err)
		return
	}

	if err = json.NewDecoder(surahBangla).Decode(&all.BanglaTranslation); err != nil {
		fmt.Println(err)
		return
	}

	surahTemplate, err := template.ParseFiles("static/html/surah.html", "static/css/sura-s.css", "static/html/common.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	// for id number being some number offset!
	offset := all.Aarabic.Verses[0].Id - 1
	for i := 0; i < len(all.Aarabic.Verses); i++ {
		all.Aarabic.Verses[i].Id = all.Aarabic.Verses[i].Id - offset
	}

	surahBuff := new(bytes.Buffer)
	surahTemplate.Execute(surahBuff, all)

	// Chashing
	SurahCash[idx] = surahBuff.Bytes()

	w.Write(surahBuff.Bytes())
}
