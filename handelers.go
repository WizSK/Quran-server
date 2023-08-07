package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	if len(r.URL.Path) == 1 {
		m := getIndex(w, r, "/")
		printStat(r, t, m)
		return
	}

	url := string(r.URL.Path[1:])
	m := getSurah(w, r, url)
	printStat(r, t, m)

}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	// id, err := surahNumCheck(string(r.URL.Path[len("/word/"):]))
	id, err := surahNumCheck(r.URL.Path[len("/w/"):])
	if id == "" {
		m := getIndex(w, r, "/w/")
		printStat(r, t, m)
		return
	}
	if err != nil {
		w.Write([]byte("<h1>page not found</h1>"))
		return
	}
	lang := "bangla"
	m := wordByWord(w, r, id, lang)
	printStat(r, t, m)
}

func wordTHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	// id, err := surahNumCheck(string(r.URL.Path[len("/word/"):]))
	id, err := surahNumCheck(r.URL.Path[len("/t/"):])
	if id == "" {
		m := getIndex(w, r, "/t/")
		printStat(r, t, m)
		return
	}
	if err != nil {
		w.Write([]byte("<h1>page not found</h1>"))
		return
	}
	lang := "bangla"
	m := wordByWordTranslation(w, r, id, lang)
	printStat(r, t, m)
}

func reDirectToWord(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	path := "/w/" + r.URL.Path[len("/word/"):]
	http.Redirect(w, r, path, http.StatusMovedPermanently)
	printStat(r, t, "rDir")
}

func surahNumCheck(index string) (string, error) {
	id, err := strconv.Atoi(index)
	if err != nil || id < 1 || id > 144 {
		return "", fmt.Errorf("index out of bound")
	}
	return index, nil
}
