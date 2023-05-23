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

type CompleteSurahWordByWord struct {

	Aarabic    Ayas
	Translaions []TranslatedVerses
	BanglaTranslation TranslatedVerses // local lang of mine.
	// Aarabic           Ayas
	// Translaions       []TranslatedVerses
	// BanglaTranslation TranslatedVerses // local lang of mine.
	SurahInfo       ChapterInfo
	WordByWordArray []WordByWord
}

type WordByWord struct {
	Verses []WordVerse
}

type WordVerse struct {
	Id       int
	VerseKey string `json:"verse_key"`
	Words    []Word
}

type Word struct {
	Position    int
	Text        string
	Translation Translation
}

type Translation struct {
	Text string
}

func pageNumbers(path string) int {
	f, err := os.ReadFile(path + "page_count.txt")
	if err != nil {
		fmt.Println(err)
		return -1
	}
	id := string(f)
	if id[len(id)-1] == '\n' {
		id = id[:len(id)-1]
	}
	pages, err := strconv.Atoi(string(id))
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return pages
}

// 'https://api.quran.com/api/v4/quran/verses/words=true&word_translation_language=bn'
func wordByWord(w http.ResponseWriter, r *http.Request, id, lang string) {
	var V []WordByWord

	path := "static/json/word_by_word/" + lang + "/" + id + "/"
	pages := pageNumbers(path)
	for i := 1; i <= pages; i++ {
		f := fmt.Sprintf("%s%d.json", path, i)
		file, err := os.Open(f)
		if err != nil {
			fmt.Println(err)
			return
		}

		var word WordByWord
		if err = json.NewDecoder(file).Decode(&word); err != nil {
			fmt.Println(err)
			return
		}
		V = append(V, word)
		file.Close()
	}

	// chapter info
	surahInfo, err := os.Open("static/json/chapters/" + id + ".json")

	if err != nil {
		w.Write([]byte("Page Not found"))
		fmt.Println(err)
		return
	}
	defer surahInfo.Close()

	var combined CompleteSurahWordByWord 
	if err = json.NewDecoder(surahInfo).Decode(&combined.SurahInfo); err != nil {
		fmt.Println(err)
		return
	}

	offset := V[0].Verses[0].Id - 1
	for i := range V {
		for j := range V[i].Verses {
			V[i].Verses[j].Id -= offset
		}

	}

	combined.WordByWordArray = V
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
