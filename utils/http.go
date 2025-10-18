package utils

import "github.com/gin-gonic/gin"

func GinHttpSuccessResponse(httpstatusCode int, message string, result map[string]any, c *gin.Context) {
	GinHttpResponse(httpstatusCode, "status", message, result, c)
}

func GinHttpFailedResponse(httpstatusCode int, message string, result map[string]any, c *gin.Context) {
	GinHttpResponse(httpstatusCode, "failed", message, result, c)
}

func GinHttpResponse(httpstatusCode int, status string, message string, result map[string]any, c *gin.Context) {
	c.JSON(httpstatusCode, gin.H{"status": status, "message": message, "result": result})
}
