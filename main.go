package main

import (
	"github.com/gin-gonic/gin"
	"go-notes/config"
	"go-notes/routes"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	routes.NoteRoutes(r)
	r.Run(":8080")
}
