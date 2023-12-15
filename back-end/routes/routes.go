package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello there!!")

	})
	r.GET("/ticks/all", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "All the ticks")
	})
	r.GET("/ticks", func(ctx *gin.Context) { //?company=Test
		companyName := ctx.DefaultQuery("company", "unknown")
		ctx.String(http.StatusOK, "Company name provided %s", companyName)
	})
	return r
}
