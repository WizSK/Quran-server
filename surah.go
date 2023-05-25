package main

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

// Cashing thing..
var SurahCash = stringByteMap()

func getSurah(w http.ResponseWriter, r *http.Request, idx string) string {

	idx, err := surahNumCheck(idx)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return "err"
	}

	if val, ok := SurahCash[idx]; ok {
		w.Write(val)
		return "cache"
	}

	var combined CompleteSurah

	if combined.SurahInfo, err = GetSurahInfo(idx); err != nil {
		fmt.Println(err)
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return "err"
	}

	if combined.Aarabic, err = GetArabicAyas(idx); err != nil {
		fmt.Println(err)
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return "err"
	}

	if combined.Translaions, err = GetTransLations(idx); err != nil {
		fmt.Println(err)
		w.Write([]byte("<h1>page not found surah number is wrong</h1>"))
		return "err"
	}

	if combined.BanglaTranslation, err = GetBanglaTranslation(idx); err != nil {
		fmt.Println(err)
		w.Write([]byte("<h1>page not found surah number is wrong</h1>"))
		return "err"
	}

	surahTemplate, err := template.ParseFiles(StaticDir+"/html/surah.html",
		StaticDir+"/css/sura-s.css",
		StaticDir+"/html/common.html",
		StaticDir+"/js/common.js",
	)

	if err != nil {
		fmt.Println(err)
		return "err"
	}

	surahBuff := new(bytes.Buffer)
	surahTemplate.Execute(surahBuff, combined)

	// Chashing
	if Cache {
		SurahCash[idx] = surahBuff.Bytes()
	}

	w.Write(surahBuff.Bytes())
	return "comp"
}
