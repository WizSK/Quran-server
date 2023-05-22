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

	st := new(bytes.Buffer)

	h, _ := template.New("html").Parse(s)
	_, err := h.ParseFiles("static/css/wordByWord.css")
	if err != nil {
		fmt.Println(err)
	}
	h.Execute(st, V)

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
	{{ range . }}

	{{ range .Verses }}
		<div class="aya">

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
		<br>
		<br>
		<br>
		<br>
		<br>
		<br>
		<br>
	{{ end }}

	{{ end }}
</body>
</html>
`
