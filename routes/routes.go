package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/atsanalyzer/handler"
)

func MountHTTPRoutes(r *gin.Engine) {

	fileHanlder := handler.NewFileHandler()
	r.POST("/analyse", fileHanlder.Analyse)
}
	