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
	check := g.Group("/api/check")
	{
		check.GET("/health", handler.HealthCheck)
		check.GET("/disk", handler.DiskCheck)
		check.GET("/cpu", handler.CPUCheck)
		check.GET("/ram", handler.RAMCheck)
	}

	// demo handler
	demo := g.Group("/api/demo")
	{
		demo.GET("/queryList", handler.QueryDemoList)
		demo.GET("/queryById/:id", handler.QueryDemoById)
		demo.POST("/add", handler.AddDemo)
		demo.POST("/update", handler.UpdateDemo)
		demo.DELETE("/delete/:id", handler.DeleteDemoById)
	}

	// account handler
	account := g.Group("/api/account/user")
	{
		account.POST("/register", handler.RegisterUserAccount)
		account.POST("/login", handler.LoginUserAccount)
	}

	return g
}

