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
	h, err := template.New("html").Parse(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = h.ParseFiles("static/js/common.js", "static/css/wordByWord.css")
	if err != nil {
		fmt.Println(err)
		return
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
  <button id="theme-tgl">Theme</button>
  <button id="fontPlus">Font +</button>
  <button id="fontMinus">Font -</button>
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
	<script>
	{{ template "theme-js" }}
	{{ template "font-size-js" }}
	</script>
</body>
</html>
`
