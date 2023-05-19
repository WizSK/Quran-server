package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", staticHandeler)

	var err error
	port := ":8080"
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
		fmt.Printf("Running at: http://localhost%s/ \n\n", port)
		err = http.ListenAndServe(port, nil)
	} else {
		fmt.Printf("Running at: http://localhost%s/ \n\n", port)
		err = http.ListenAndServe(port, nil)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	if len(r.URL.Path) == 1 {
		getIndex(w, r)
		printStat(r, t)
		return
	}

	url := string(r.URL.Path[1:])
	getSurah(w, r, url)
	printStat(r, t)

}

func printStat(r *http.Request, dur time.Time) {
	fmt.Printf("[stat] %s | %13s | %15s | %s | \"%s\"\n",
		time.Now().Format("2006/01/02 - 03:04:05 PM"),
		time.Since(dur),
		strings.Split(r.RemoteAddr, ":")[0],
		r.Method,
		r.URL.Path,
	)
}
