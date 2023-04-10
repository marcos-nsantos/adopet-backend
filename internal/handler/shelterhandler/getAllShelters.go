package shelterhandler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

// GetAllShelters handles request to get all shelters
//
//	@Summary	Get all shelters
//	@Tags		shelter
//	@Produce	json
//	@Param		page	query		int	false	"Page number"					default(1)
//	@Param		limit	query		int	false	"Limit of shelters per page"	default(10)
//	@Success	200		{object}	schemas.SheltersResponse
//	@Router		/shelters [get]
func GetAllShelters(c *gin.Context) {
	page, limit := queryShelters(c)

	shelters, total, err := database.GetAllShelters(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting shelters"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToSheltersResponse(shelters, page, limit, total))
}

func queryShelters(c *gin.Context) (int, int) {
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
