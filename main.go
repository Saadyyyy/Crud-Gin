package main

import (
	c "Crud-Gin/app/config"
	"Crud-Gin/app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	db := c.InitializeDatabase()

	routes.Api(r, db)
	routes.ImagesRouter(r, db)
	r.Run()
}
