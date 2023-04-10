package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/docs"
	"github.com/marcos-nsantos/adopet-backend/internal/auth"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/adoptionhandler"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/authhandler"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/pethandler"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/shelterhandler"
	"github.com/marcos-nsantos/adopet-backend/internal/handler/tutorhandler"
	swaggerfiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	authentication := r.Group("/auth")
	{
		authentication.POST("/tutor", authhandler.AuthenticateTutor)
		authentication.POST("/shelter", authhandler.AuthenticateShelter)
	}

	tutors := r.Group("/tutors")
	{
		tutors.Use(auth.JWTAuth())

		tutors.GET("/:id", tutorhandler.GetTutorByID)
		tutors.GET("", tutorhandler.GetAllTutors)
		tutors.PUT("/:id", tutorhandler.UpdateTutor)
		tutors.DELETE("/:id", tutorhandler.DeleteTutor)
	}

	shelters := r.Group("/shelters")
	{
		shelters.Use(auth.JWTAuth())

		shelters.GET("/:id", shelterhandler.GetShelterByID)
		shelters.GET("", shelterhandler.GetAllShelters)
		shelters.PUT("/:id", shelterhandler.UpdateShelter)
		shelters.DELETE("/:id", shelterhandler.DeleteShelter)
	}

	pets := r.Group("/pets")
	{
		pets.Use(auth.JWTAuth())

		pets.POST("", pethandler.CreatePet)
		pets.GET("/:id", pethandler.GetPetByID)
		pets.GET("", pethandler.GetAllPets)
		pets.PUT("/:id", pethandler.UpdatePet)
		pets.PATCH("/:id/adopted", pethandler.UpdateIsAdoptedPet)
		pets.DELETE("/:id", pethandler.DeletePet)
	}

	adoption := r.Group("/adoptions")
	{
		adoption.Use(auth.JWTAuth())

		adoption.POST("/:petId/:tutorId", adoptionhandler.CreateAdoption)
		adoption.DELETE("/:id", adoptionhandler.DeleteAdoption)
	}

	r.POST("tutors", tutorhandler.CreateTutor)
	r.POST("shelters", shelterhandler.CreateShelter)

	docs.SwaggerInfo.Title = "Adopet API"
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerfiles.Handler))

	return r
}
