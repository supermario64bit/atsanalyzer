package types

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/atsanalyzer/utils"
)

type ApplicationError struct {
	HttpStatusCode int
	Message        string
	Error          error
}

func (appErr *ApplicationError) GinHttpResponse(c *gin.Context) {
	utils.GinHttpFailedResponse(appErr.HttpStatusCode, appErr.Message, gin.H{"error": appErr.Error.Error()}, c)
}
