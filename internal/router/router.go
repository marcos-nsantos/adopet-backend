package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/pethandler"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/shelterhandler"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/tutorhandler"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	tutors := r.Group("/tutors")
	{
		tutors.POST("", tutorhandler.CreateTutor)
		tutors.GET("/:id", tutorhandler.GetTutorByID)
		tutors.GET("", tutorhandler.GetAllTutors)
		tutors.PUT("/:id", tutorhandler.UpdateTutor)
		tutors.DELETE("/:id", tutorhandler.DeleteTutor)
	}

	shelters := r.Group("/shelters")
	{
		shelters.POST("", shelterhandler.CreateShelter)
		shelters.GET("/:id", shelterhandler.GetShelterByID)
		shelters.GET("", shelterhandler.GetAllShelters)
		shelters.PUT("/:id", shelterhandler.UpdateShelter)
		shelters.DELETE("/:id", shelterhandler.DeleteShelter)
	}

	pets := r.Group("/pets")
	{
		pets.POST("", pethandler.CreatePet)
		pets.GET("/:id", pethandler.GetPetByID)
		pets.GET("", pethandler.GetAllPets)
		pets.PUT("/:id", pethandler.UpdatePet)
	}

	return r
}
