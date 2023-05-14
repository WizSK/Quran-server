package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	prot := "8000"
	if len(os.Args) == 2 {
		prot = os.Args[1]
	}
	route := gin.Default()

	route.GET("/", getIndex)
	route.GET("/:id", getSurah)

	route.StaticFile("static/fonts/arabic", "static/assets/uthman_tn09.otf")
	route.StaticFile("static/fonts/bangla", "static/assets/SolaimanLipi.ttf")
	route.StaticFile("static/fonts/english", "static/assets/Lato-Regular.ttf")
	route.StaticFile("static/images/favicon", "static/assets/quran-faviocn.png")
	route.StaticFile("static/images/quran.png", "static/assets/quran.png")

	route.Run(":" + prot)
}
