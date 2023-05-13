package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/", getIndex)
	route.GET("/:id", getSurah)

	route.StaticFile("static/fonts/arabic", "static/assets/uthman_tn09.otf")
	route.StaticFile("static/fonts/bangla", "static/assets/SolaimanLipi.ttf")
	route.StaticFile("static/fonts/english", "static/assets/Lato-Regular.ttf")
	route.StaticFile("static/images/favicon", "static/assets/quran-faviocn.png")
	route.StaticFile("static/images/quran.png", "static/assets/quran.png")

	route.Run(":8000")
}
