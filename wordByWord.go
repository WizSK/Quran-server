package main

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

const WordCache bool = true

var WordByWordCache = func() map[string][]byte {
	if WordCache {
		return make(map[string][]byte)

	}
	return nil
}()

func wordByWord(w http.ResponseWriter, r *http.Request, index, lang string) string {
	if WordCache {
		if val, ok := WordByWordCache[index]; ok {
			w.Write(val)
			return "cache"
		}
	}
	var combined CompleteSurahWordByWord
	var err error

	if combined.SurahInfo, err = GetSurahInfo(index); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return "err"
	}

	if combined.WordByWordArray, err = GetWordByWord(index, lang); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return "err"
	}

	wordHmtl, err := template.ParseFiles("static/html/word.html",
		"static/html/common.html",
		"static/js/common.js",
		"static/css/wordByWord.css",
	)
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	st := new(bytes.Buffer)
	if err = wordHmtl.Execute(st, combined); err != nil {
		fmt.Println(err)
		return "err"
	}
	if WordCache {
		WordByWordCache[index] = st.Bytes()
	}
	w.Write(st.Bytes())
	return "comp"
}
