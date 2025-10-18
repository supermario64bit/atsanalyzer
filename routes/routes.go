package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/atsanalyzer/handler"
)

func MountHTTPRoutes(r *gin.Engine) {

	resumeAnalysisHandler := handler.NewResumeAnalysisHandler()
	r.POST("/analyse", resumeAnalysisHandler.Analyse)
}
