package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	router.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
