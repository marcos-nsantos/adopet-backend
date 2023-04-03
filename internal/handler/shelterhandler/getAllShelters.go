package shelterhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
	"github.com/marcos-nsantos/adopet-backend/internal/schemas"
)

func GetAllShelters(c *gin.Context) {
	page, limit := queryShelters(c)

	shelters, total, err := database.GetAllShelters(page, limit)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting shelters"})
		return
	}

	c.JSON(http.StatusOK, schemas.ToUsersResponse(shelters, page, limit, total))
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
