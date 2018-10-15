package router


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dapphouse/router/middleware"
	"dapphouse/handler"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The Incorrect API Route")
	})

	// the health check handler
	checkGin := g.Group("api/check")
	{
		checkGin.GET("/health", handler.HealthCheck)
		checkGin.GET("/disk", handler.DiskCheck)
		checkGin.GET("/cpu", handler.CPUCheck)
		checkGin.GET("/ram", handler.RAMCheck)
	}

	// demo handler
	demoGin := g.Group("/api/demo")
	{
		demoGin.GET("/queryList", handler.QueryDemoList)
		demoGin.GET("/queryById/:id", handler.QueryDemoById)
		demoGin.POST("/add", handler.AddDemo)
		demoGin.POST("/update", handler.UpdateDemo)
		demoGin.DELETE("/delete/:id", handler.DeleteDemoById)

	}

	return g
}

