package router

import (
	"api-server/handler/sd"
	"api-server/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c*gin.Context) {
		c.String(http.StatusNotFound, "the incorrect API route.")
	})

	healthGroup:= g.Group("/sd")
	{
		healthGroup.GET("/health", sd.HealthCheck)
		healthGroup.GET("/disk", sd.DiskCheck)
		healthGroup.GET("/cpu", sd.CPUCheck)
		healthGroup.GET("/ram", sd.RAMCheck)
	}
}
