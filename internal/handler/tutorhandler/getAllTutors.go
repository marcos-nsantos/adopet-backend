package tutorhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/marcos-nsantos/adopet-backend/internal/schemas"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

// GetAllTutors handles request to get all tutors
//
//	@Summary	Get all tutors
//	@Tags		tutors
//	@Security	Bearer
//	@Produce	json
//	@Param		page	query		int	false	"Page number"				default(1)
//	@Param		limit	query		int	false	"Limit of tutors per page"	default(10)
//	@Success	200		{object}	schemas.TutorsResponse
//	@Router		/tutors [get]
func GetAllTutors(c *gin.Context) {
	page, limit := queryTutors(c)

	tutors, total, err := database.GetAllTutors(page, limit)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting tutors"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToTutorsResponses(tutors, page, limit, total))
}

func queryTutors(c *gin.Context) (int, int) {
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
