package main

import (
	"fmt"
	"log"
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
	router.GET("/dates/:year", SearchYear)

	router.Run("localhost:8080")
}

func getDate(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dates)
}

func HandlerFunc(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		log.Println(err)
		return
	}
	result := SearchYear(year)
	c.String(http.StatusOK, result)
}

func SearchYear(year int) string {
	now := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	if year < time.Now().Year() {
		daysGone := int64(time.Since(now).Hours()) / 24
		return fmt.Sprintf("Days gone: %d", daysGone)
	} else {
		daysLeft := int64(time.Until(now).Hours()) / 24
		return fmt.Sprintf("Days left: %d", daysLeft)
	}
}
