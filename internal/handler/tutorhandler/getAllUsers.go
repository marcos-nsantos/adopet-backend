package tutorhandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marcos-nsantos/adopet-backend/internal/database"
)

func GetAllTutors(c *gin.Context) {
	page, limit := queryTutors(c)

	tutors, total, err := database.GetAllTutors(page, limit)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting tutors"})
		return
	}

	c.JSON(http.StatusOK, toTutorsResponse(tutors, page, limit, total))
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
