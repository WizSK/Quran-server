package main

import (
	"net/http"
	"strings"
	"time"
)

/*
"static/fonts/arabic", "static/assets/uthman_tn09.otf")
"static/fonts/bangla", "static/assets/SolaimanLipi.ttf")
"static/fonts/english", "static/assets/Lato-Regular.ttf")
"static/images/favicon", "static/assets/quran-faviocn.png")
"static/images/quran.png", "static/assets/quran.png")
*/

func staticHandeler(w http.ResponseWriter, r *http.Request) {
	dur := time.Now()
	v := strings.Split(r.URL.Path[1:], "/")

	if len(v) < 2 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no file"))
		return
	}

	if v[1] == "fonts" {
		if len(v) < 3 || len(v) > 3 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("no file"))
			return
		}

		switch v[2] {
		case "arabic":
			http.ServeFile(w, r, "static/assets/uthman_tn09.otf")
			printStat(w, r, dur)
			return

		case "english":
			http.ServeFile(w, r, "static/assets/Lato-Regular.ttf")
			printStat(w, r, dur)
			return

		case "bangla":
			http.ServeFile(w, r, "static/assets/SolaimanLipi.ttf")
			printStat(w, r, dur)
			return

		default:
			w.Write([]byte("no file found"))
			printStat(w, r, dur)
			return
		}

	}

	if v[1] == "images" {
		if len(v) < 3 || len(v) > 3 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("no file"))
			return
		}
		switch v[2] {
		case "favicon":
			http.ServeFile(w, r, "static/assets/quran-faviocn.png")
			printStat(w, r, dur)
			return

		case "quran.png":
			http.ServeFile(w, r, "static/assets/quran.png")
			printStat(w, r, dur)
			return

		default:
			w.Write([]byte("no file found"))
			printStat(w, r, dur)
			return
		}

	}
}
