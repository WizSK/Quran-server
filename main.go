package main

import (
	"fmt"
	"net/http"
	"os"
)

// var Cache bool = QuranCacheEnv()
var Cache bool = false
const StaticDir string = "static/"

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/static/", staticHandeler)
	http.HandleFunc("/word/", reDirectToWord)
	http.HandleFunc("/w/", wordHandler)
	http.HandleFunc("/t/", wordTHandler)

	var err error
	port := ":8001"
	if len(os.Args) == 2 {
		port = ":" + os.Args[1]
		fmt.Printf("Running at: http://localhost%s/ \n\n", port)
		err = http.ListenAndServe(port, nil)
	} else {
		fmt.Printf("Running at: http://localhost%s/\nOr on your docker specified port \n\n", port)
		err = http.ListenAndServe(port, nil)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func QuranCacheEnv() bool {
	c := os.Getenv("QURAN_CACHE")
	if c == "" {
		fmt.Printf("\nCACHE = False\n")
		return false
	}
	fmt.Printf("\nCACHE = True\n")
	return true
}

func StaticDirEnv() string {
	c := os.Getenv("STATIC_DIR")
	if c == "" {
		fmt.Printf("static file dir `static`\n\n")
		return "static"
	}
	fmt.Printf("static file dir `%s`\n\n", c)
	return c
}
