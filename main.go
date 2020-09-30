package main

import (
	"github.com/1000Delta/UserAuth/config"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	cfg := config.DefaultConfig()
	
	gin.SetMode(cfg.Mode)
	RegisterRoutes(g)

	g.Run(cfg.Listen)
}
