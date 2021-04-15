package services

import "github.com/gin-gonic/gin"

func DecideDate(c *gin.Context) (string, string) {
	queryParams := c.Request.URL.Query()

	var decidedStartDate string
	var decidedEndDate string

	if len(queryParams["start_date"]) == 0 {
		decidedStartDate = "1990-01-01"
	} else {
		decidedStartDate = queryParams["start_date"][0]
	}

	if len(queryParams["end_date"]) == 0 {
		decidedEndDate = "4000-12-12"
	} else {
		decidedEndDate = queryParams["end_date"][0]
	}

	return decidedStartDate, decidedEndDate
}
