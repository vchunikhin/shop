package pkg

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"message"`
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	log.Error(message)
	ctx.AbortWithStatusJSON(statusCode, Error{message})
}
