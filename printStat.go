package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// print the pretty output ^_^
func printStat(r *http.Request, dur time.Time, method string) {
	reset := "\033[0m"
	// default is green maing cache
	methodBg := "\033[42m"
	// Black
	methodFg := "\033[0;30m"

	if method == "comp" {
		// yellow
		methodBg = "\033[43m"
	}
	if method == "err" {
		// red
		methodBg = "\033[41m"
	}

	fmt.Printf("[stat] %s | %13s |%s%s %5s %s| %15s | %s | \"%s\"\n",
		time.Now().Format("2006/01/02 - 03:04:05 PM"),
		time.Since(dur),
		methodFg,
		methodBg,
		method,
		reset,
		strings.Split(r.RemoteAddr, ":")[0],
		r.Method,
		r.URL.Path,
	)
}

/*
// laaaaaaaaaaaaaaaa
func colorableStdout() io.Writer {
		if runtime.GOOS == "windows" {
			return colorable.NewColorableStdout()
		}
	return os.Stdout
}
*/
