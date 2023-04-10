package pethandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// GetAllPets handles request to get all pets
//
//	@Summary	Get all pets
//	@Tags		pets
//	@Security	Bearer
//	@Param		page	query	int	false	"Page number"				default(1)
//	@Param		limit	query	int	false	"Number of pets per page"	default(10)
//	@Produce	json
//	@Success	200	{object}	schemas.PetsResponse
//	@Router		/pets [get]
func GetAllPets(c *gin.Context) {
	page, limit := queryPets(c)

	pets, total, err := database.GetAllPets(page, limit)
	if err != nil {
		log.Println("error to get pets", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to get pets"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToPetsResponse(pets, page, limit, total))
}

func queryPets(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}

	return page, limit
}
