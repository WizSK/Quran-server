package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var IndexCash []byte
var WordByWordIndexCash []byte

func getIndex(w http.ResponseWriter, r *http.Request, prefix string)string {
	// Cashed
	switch prefix {
	case "/":
		if len(IndexCash) > 0 {
			w.Write(IndexCash)
			return "cache"
		}
	case "/w/":
		if len(WordByWordIndexCash) > 0 {
			w.Write(WordByWordIndexCash)
			return "cache"
		}

	}

	// const surahUrl string = "https://api.quran.com/api/v4/chapters"
	suras := new(bytes.Buffer)
	// resp, err := http.Get(url)
	resp, err := os.Open("static/json/chapters.json")
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	// if resp.StatusCode != http.StatusOK {
	// 	resp.Body.Close()
	// 	return suras, fmt.Errorf("http response: %d from %s", resp.StatusCode, url)
	// }
	defer resp.Close()

	var surahJson ChaptersIdx
	if err = json.NewDecoder(resp).Decode(&surahJson); err != nil {
		fmt.Println(err)
		return "err"
	}

	var prefixedSurah SurahIndexPrefixed
	for _, v := range surahJson.Chapters {
		p := fmt.Sprintf("%s%d", prefix, v.Id)
		prefixedSurah.Prefixes = append(prefixedSurah.Prefixes, p)
	}
	prefixedSurah.ChaptersIdx = surahJson

	// prefixedSurah.ChaptersIdx.Chapters[0].Id
	// prefixedSurah.Prefixes[0]

	p, err := template.ParseFiles("static/html/index.html", "static/css/index.css", "static/html/common.html")
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	p.Execute(suras, prefixedSurah)

	switch prefix {
	case "/":
		IndexCash = suras.Bytes()
	case "/w/":
		WordByWordIndexCash = suras.Bytes()
	}
	// IndexCash = suras.Bytes()
	w.Write(suras.Bytes())
		return "comp"
}
