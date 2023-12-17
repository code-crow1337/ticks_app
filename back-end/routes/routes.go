package routes

import (
	"e-identitet/tick-api/data"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	logger := log.Default()
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	db := data.InitDatabase()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello there!!")

	})
	r.GET("/company", func(ctx *gin.Context) {
		companyNames := GetCompanyNames(db)
		logger.Print("🥑", companyNames)
		ctx.JSON(http.StatusOK, companyNames)

	})
	r.GET("/company/:name", func(ctx *gin.Context) { //single company name
		companyName := ctx.Param("name")
		ctx.JSON(http.StatusOK, companyName)
	})
	r.GET("/tick", func(ctx *gin.Context) { //maybe not needed
		ctx.String(http.StatusOK, "All the ticks")
	})
	r.GET("/tick/:tickType", func(ctx *gin.Context) {
		tickType := ctx.Param("tickType")
		ctx.JSON(http.StatusOK, tickType)
	})
	return r
}
