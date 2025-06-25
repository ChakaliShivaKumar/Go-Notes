package routes

import (
	"github.com/gin-gonic/gin"
	"go-notes/controllers"
)

func NoteRoutes(r *gin.Engine) {
	r.POST("/notes", controllers.CreateNote)
	r.GET("/notes", controllers.GetNotes)
	r.GET("/notes/:id", controllers.GetNoteByID)
	r.PUT("/notes/:id", controllers.UpdateNote)
	r.DELETE("/notes/:id", controllers.DeleteNote)
}
