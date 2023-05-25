package main

import (
	"net/http"
	"strings"
	"time"
)

func staticHandeler(w http.ResponseWriter, r *http.Request) {
	dur := time.Now()
	v := strings.Split(r.URL.Path[1:], "/")

	if len(v) < 2 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no file"))
		printStat(r, dur, "err")
		return
	}

	if v[1] == "fonts" {
		if len(v) < 3 || len(v) > 3 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("no file"))
			printStat(r, dur, "err")
			return
		}

		switch v[2] {
		case "arabic":
			http.ServeFile(w, r, StaticDir+"/assets/UthmanicHafs1Ver18.woff2")
			printStat(r, dur, "file")
			return

		case "english":
			http.ServeFile(w, r, StaticDir+"/assets/Lato-Regular.ttf")
			printStat(r, dur, "file")
			return

		case "bangla":
			http.ServeFile(w, r, StaticDir+"/assets/SolaimanLipi.ttf")
			printStat(r, dur, "file")
			return

		default:
			w.Write([]byte("no file found"))
			printStat(r, dur, "err")
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
			http.ServeFile(w, r, StaticDir+"/assets/quran-faviocn.png")
			printStat(r, dur, "file")
			return

		case "quran.png":
			http.ServeFile(w, r, StaticDir+"/assets/quran.png")
			printStat(r, dur, "file")
			return

		default:
			w.Write([]byte("no file found"))
			printStat(r, dur, "err")
			return
		}

	}
}
