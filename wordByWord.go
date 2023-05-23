package main

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"
)

// 'https://api.quran.com/api/v4/quran/verses/words=true&word_translation_language=bn'
func wordByWord(w http.ResponseWriter, r *http.Request, index, lang string) {
	var combined CompleteSurahWordByWord
	var err error

	if combined.SurahInfo, err = GetSurahInfo(index); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return
	}

	if combined.WordByWordArray, err = GetWordByWord(index, lang); err != nil {
		w.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return
	}

	st := new(bytes.Buffer)
	h, _ := template.New("html").Parse(s)
	_, err = h.ParseFiles("static/css/wordByWord.css")
	if err != nil {
		fmt.Println(err)
	}
	h.Execute(st, combined)

	w.Write(st.Bytes())
}

const s string = `<!DOCTYPE html lang="en">
<html>
<head>
	<style>
	{{ template "wordByWord" }}
	</style>
</head>
<body>
<section id="hero">
<h1>{{ .SurahInfo.Chapter.Name }}</h1>
<div id="place">{{ .SurahInfo.Chapter.Place }}</div>
{{ if .SurahInfo.Chapter.Bismillah }}
    <p id="bismillah">بِسْمِ ٱللَّهِ ٱلرَّحْمَـٰنِ ٱلرَّحِيمِ</p>
{{ end }}
</section>


{{ range .WordByWordArray }}

{{ range .Verses }}
	<div class="aya" id="{{ .Id }}">

	<div class="verse_key">{{ .VerseKey }}</div>
	<div class="words">
	{{ range .Words }}
	<span class="word">
		<span class="arabic">
		{{ .Text }}
		</span>
		<span class="trans">
		{{ .Translation.Text }}
		</span>
	</span>
	{{ end }}
	</div>
	</div>
{{ end }}

{{ end }}
</body>
</html>
`
