package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type date struct {
	year int `json:"year"`
}

var dates = []date{
	{year: 2022},
	{year: 2000},
	{year: 2005},
}

func main() {
	router := gin.Default()
	router.GET("/dates", getDate)
	router.GET("/dates/:year", getDaysByYear)

	router.Run("localhost:8080")
}

func getDate(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dates)
}

func getDaysByYear(c *gin.Context) {
	t := time.Now().Year()

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong year"})
		return
	}

	if year < t {
		c.IndentedJSON(http.StatusOK, gin.H{"Days gone": (t - year) * 365})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"Days left": (year - t) * 365})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong year"})
}
