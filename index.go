package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var IndexCache = stringByteMap()

func stringByteMap() map[string][]byte {
	if Cache {
		return make(map[string][]byte)
	}
	return nil
}

func getIndex(w http.ResponseWriter, r *http.Request, prefix string) string {
	// Cashed
	if val, ok := IndexCache[prefix]; ok {
		w.Write(val)
		return "cache"
	}

	// const surahUrl string = "https://api.quran.com/api/v4/chapters"
	suras := new(bytes.Buffer)
	// resp, err := http.Get(url)
	resp, err := os.Open(StaticDir + "/json/chapters.json")
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	// if resp.StatusCode != http.StatusOK {
	// 	resp.Body.Close()
	// 	return suras, fmt.Errorf("http response: %d from %s", resp.StatusCode, url)
	// }
	defer resp.Close()

	var surahJson ChaptersIdx
	if err = json.NewDecoder(resp).Decode(&surahJson); err != nil {
		fmt.Println(err)
		return "err"
	}

	var prefixedSurah SurahIndexPrefixed
	for _, v := range surahJson.Chapters {
		p := fmt.Sprintf("%s%d", prefix, v.Id)
		prefixedSurah.Prefixes = append(prefixedSurah.Prefixes, p)
	}
	prefixedSurah.ChaptersIdx = surahJson

	// prefixedSurah.ChaptersIdx.Chapters[0].Id
	// prefixedSurah.Prefixes[0]

	p, err := template.ParseFiles(StaticDir+"/html/index.html", StaticDir+"/css/index.css", StaticDir+"/html/common.html", StaticDir+"/js/common.js")
	if err != nil {
		fmt.Println(err)
		return "err"
	}

	p.Execute(suras, prefixedSurah)

	if IndexCache != nil {
		IndexCache[prefix] = suras.Bytes()
	}
	// IndexCash = suras.Bytes()
	w.Write(suras.Bytes())
	return "comp"
}
