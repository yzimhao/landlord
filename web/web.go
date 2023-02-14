package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	www := gin.New()
	www.LoadHTMLGlob("./web/template/default/*.html")
	www.StaticFS("/statics", http.Dir("./web/template/default/statics"))

	www.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	www.Run(":8081")
}
