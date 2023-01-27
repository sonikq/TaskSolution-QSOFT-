package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorRespose struct {
	Message string `json: "message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorRespose{message})
}

func ChangeHeader(c *gin.Context) {
	header := c.GetHeader("X-PING")
	if header == "PING" {
		c.Header().Add("X-PONG", "PONG")
		return
	}

}
