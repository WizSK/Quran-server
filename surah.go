package main

import (
	"bytes"
	"fmt"
	"net/http"
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

	var combined CompleteSurah

	if combined.SurahInfo, err = GetSurahInfo(idx); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return
	}

	if combined.Aarabic, err = GetArabicAyas(idx); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return
	}

	if combined.Translaions, err = GetTransLations(idx); err != nil {
		w.Write([]byte("<h1>page not found surah number is wrong</h1>"))
		return
	}

	if combined.BanglaTranslation, err = GetBanglaTranslation(idx);err!=nil {
		w.Write([]byte("<h1>page not found surah number is wrong</h1>"))
		return
	}


	surahTemplate, err := template.ParseFiles("static/html/surah.html", "static/css/sura-s.css", "static/html/common.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	surahBuff := new(bytes.Buffer)
	surahTemplate.Execute(surahBuff, combined)

	// Chashing
	SurahCash[idx] = surahBuff.Bytes()

	w.Write(surahBuff.Bytes())
}
