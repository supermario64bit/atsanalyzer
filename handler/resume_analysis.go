package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/service"
)

type ResumeAnalysisHandler interface {
	Analyse(c *gin.Context)
}

type resumeAnalysisHandler struct {
	svc service.ResumeAnalysis
}

func NewResumeAnalysisHandler() ResumeAnalysisHandler {
	return &resumeAnalysisHandler{
		svc: service.NewResumeAnalysisService(),
	}
}

func (h *resumeAnalysisHandler) Analyse(c *gin.Context) {
	var dto dto.ResumeRequest

	if err := c.ShouldBind(&dto); err != nil {
		log.Println("Unable to bind request body. Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request body", "result": gin.H{"error": err.Error()}})
		return
	}

	if len(dto.JobDescription) < 20 || dto.ResumeFile == nil {
		log.Println("Empty request")
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request body", "result": gin.H{"error": "Empty request"}})
		return
	}

	response, appErr := h.svc.Analyse(&dto)
	if appErr != nil {
		appErr.GinHttpResponse(c)
		return
	}

	log.Println("ATS Score available")
	c.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "Analysis completed!!", "result": gin.H{"analysis_report": response}})
}
