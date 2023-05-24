package main

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

var WordTranByWordCache = stringByteMap()

func wordByWordTranslation(w http.ResponseWriter, r *http.Request, index, lang string) string {
	// Cache
	if val, ok := WordTranByWordCache[index]; ok {
		w.Write(val)
		return "cache"
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

	if combined.BanglaTranslation, err = GetBanglaTranslation(index); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return "err"
	}

	if combined.Translaions, err = GetTransLations(index); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return "err"
	}

	wordHmtl, err := template.ParseFiles("static/html/word_trans.html",
		"static/html/common.html",
		"static/js/common.js",
		"static/css/wordByWord.css",
	)
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	_ = combined.BanglaTranslation.Translations[0].Text
	st := new(bytes.Buffer)
	if err = wordHmtl.Execute(st, combined); err != nil {
		fmt.Println(err)
		return "err"
	}

	// Cache
	if Cache {
		WordTranByWordCache[index] = st.Bytes()
	}
	w.Write(st.Bytes())
	return "comp"
}
