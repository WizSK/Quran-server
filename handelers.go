package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	if len(r.URL.Path) == 1 {
		getIndex(w, r, "/")
		printStat(r, t)
		return
	}

	url := string(r.URL.Path[1:])
	getSurah(w, r, url)
	printStat(r, t)

}

// print the pretty output ^_^
func printStat(r *http.Request, dur time.Time) {
	fmt.Printf("[stat] %s | %13s | %15s | %s | \"%s\"\n",
		time.Now().Format("2006/01/02 - 03:04:05 PM"),
		time.Since(dur),
		strings.Split(r.RemoteAddr, ":")[0],
		r.Method,
		r.URL.Path,
	)
}

func wordHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	// id, err := surahNumCheck(string(r.URL.Path[len("/word/"):]))
	id, err := surahNumCheck(r.URL.Path[len("/w/"):])
	if id == "" {
		getIndex(w, r, "/w/")
		printStat(r, t)
		return
	}
	if err != nil {
		w.Write([]byte("<h1>page not found</h1>"))
		return
	}
	lang := "bangla"
	wordByWord(w, r, id, lang)
	printStat(r, t)
}

func reDirectToWord(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	path := "/w/" + r.URL.Path[len("/word/"):]
	http.Redirect(w, r, path, 301)
	printStat(r, t)
}

func surahNumCheck(index string) (string, error) {
	id, err := strconv.Atoi(index)
	if err != nil || id < 1 || id > 144 {
		return "", fmt.Errorf("index out of bound")
	}
	return index, nil
}
