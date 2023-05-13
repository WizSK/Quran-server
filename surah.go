package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

// Cashing thing..
var SurahCash = make(map[string][]byte)

func getSurah(c *gin.Context) {
	idx := c.Param("id")

	id, err := strconv.Atoi(idx)
	if err != nil || id < 1 || id > 144 {
		c.Writer.Write([]byte("<h1>Page not found surah number is wrong</h1>"))
		return
	}

	if _, ok := SurahCash[idx]; ok {
		c.Writer.Write(SurahCash[idx])
		return
	}

	surahArabic, err := os.Open("static/json/arabic/" + idx + ".json")
	if err != nil {
		c.Writer.Write([]byte("Page Not found"))
		fmt.Println(err)
		return
	}
	defer surahArabic.Close()

	surahInfo, err := os.Open("static/json/chapters/" + idx + ".json")

	if err != nil {
		c.Writer.Write([]byte("Page Not found"))
		fmt.Println(err)
		return
	}
	defer surahInfo.Close()

	surahEnlish, err := os.Open("static/json/english/" + idx + ".json")

	if err != nil {
		c.Writer.Write([]byte("Page Not found"))
		fmt.Println(err)
		return
	}
	defer surahEnlish.Close()

	surahBangla, err := os.Open("static/json/bangla/" + idx + ".json")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer surahBangla.Close()

	var all CompleteSurah
	// ara
	if err = json.NewDecoder(surahArabic).Decode(&all.Aarabic); err != nil {
		fmt.Println(err)
		return
	}

	// info
	if err = json.NewDecoder(surahInfo).Decode(&all.SurahInfo); err != nil {
		fmt.Println(err)
		return
	}

	// trans
	all.Tranlaions = append(all.Tranlaions, TranslatedVerses{})
	if err = json.NewDecoder(surahEnlish).Decode(&all.Tranlaions[0]); err != nil {
		fmt.Println(err)
		return
	}

	all.Tranlaions = append(all.Tranlaions, TranslatedVerses{})
	if err = json.NewDecoder(surahBangla).Decode(&all.Tranlaions[1]); err != nil {
		fmt.Println(err)
		return
	}

	surahTemplate, err := template.ParseFiles("static/html/surah.html", "static/css/sura-s.css", "static/html/common.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	// for id number being some number offset!
	offset := all.Aarabic.Verses[0].Id - 1
	for i := 0; i < len(all.Aarabic.Verses); i++ {
		all.Aarabic.Verses[i].Id = all.Aarabic.Verses[i].Id - offset
	}

	surahBuff := new(bytes.Buffer)
	surahTemplate.Execute(surahBuff, all)

	// Chashing
	SurahCash[idx] = surahBuff.Bytes()

	c.Writer.Write(surahBuff.Bytes())
}
