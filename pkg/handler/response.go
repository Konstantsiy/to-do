package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseError struct {
	Message string
}

type statusResponse struct {
	Status string
}

func NewResponseError(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, responseError{message})
}
