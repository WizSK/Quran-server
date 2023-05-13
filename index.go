package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
	"os"

	"github.com/gin-gonic/gin"
)

var IndexCash []byte

func getIndex(c *gin.Context) {

	// Cashed
	if len(IndexCash) > 0 {
		c.Writer.Write(IndexCash)
		return
	}


	// const surahUrl string = "https://api.quran.com/api/v4/chapters"
	suras := new(bytes.Buffer)
	// resp, err := http.Get(url)
	resp, err := os.Open("static/json/chapters.json")

	if err != nil {
		//
		return
	}

	// if resp.StatusCode != http.StatusOK {
	// 	resp.Body.Close()
	// 	return suras, fmt.Errorf("http response: %d from %s", resp.StatusCode, url)
	// }
	defer resp.Close()

	var surahJson ChaptersIdx
	if err = json.NewDecoder(resp).Decode(&surahJson); err != nil {
		return
	}

	p, err := template.ParseFiles("static/html/index.html", "static/css/index.css", "static/html/common.html")
	if err != nil {
		fmt.Println(err)
	}

	p.Execute(suras, surahJson)

	IndexCash = suras.Bytes()
	c.Writer.Write(suras.Bytes())
}
