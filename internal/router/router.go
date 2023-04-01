package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/tutorhandler"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	tutors := r.Group("/tutors")
	{
		tutors.POST("", tutorhandler.CreateTutor)
		tutors.GET("/:id", tutorhandler.GetTutorByID)
		tutors.GET("", tutorhandler.GetAllTutors)
	}

	return r
}
