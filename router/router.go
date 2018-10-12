package router


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"dapphouse/router/middleware"
	check "dapphouse/handler/check"
	demo "dapphouse/handler/demo"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The Incorrect API Route.")
	})

	// the health check handler
	checkGin := g.Group("api/check")
	{
		checkGin.GET("/health", check.HealthCheck)
		checkGin.GET("/disk", check.DiskCheck)
		checkGin.GET("/cpu", check.CPUCheck)
		checkGin.GET("/ram", check.RAMCheck)
	}

	// demo handler
	demoGin := g.Group("/api/demo")
	{
		demoGin.GET("/queryList", demo.QueryDemoList)
		demoGin.GET("/queryById/:id", demo.QueryDemoById)
		demoGin.POST("/add", demo.AddDemo)
		demoGin.POST("/update", demo.UpdateDemo)
		demoGin.DELETE("/delete/:id", demo.DeleteDemoById)

	}

	return g
}

