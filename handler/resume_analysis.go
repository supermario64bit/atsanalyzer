package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/jobs"
	job_routes "github.com/supermario64bit/atsanalyzer/jobs/routes"
	"github.com/supermario64bit/atsanalyzer/service"
	"github.com/supermario64bit/atsanalyzer/utils"
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
	var reqDto dto.ResumeRequest

	if err := c.ShouldBind(&reqDto); err != nil {
		log.Println("Unable to bind request body. Error: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request body", "result": gin.H{"error": err.Error()}})
		return
	}

	if len(reqDto.JobDescription) < 20 || reqDto.ResumeFile == nil {
		log.Println("Empty request")
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request body", "result": gin.H{"error": "Empty request"}})
		return
	}

	resumeText, err := utils.ExtractTextFromPDF(reqDto.ResumeFile)

	jsonData, err := json.Marshal(dto.ResumeAnalysisJob{JobDescription: reqDto.JobDescription, ResumeTest: resumeText})
	if err != nil {
		utils.GinHttpFailedResponse(http.StatusInternalServerError, "Unable to analyse resume", gin.H{"error": err.Error()}, c)
		return
	}
	task := asynq.NewTask(job_routes.ResumeAnalyseRoute, jsonData)

	asynqClient := jobs.NewAsynqClient()
	defer asynqClient.Close()
	if _, err := asynqClient.Enqueue(task); err != nil {
		utils.GinHttpFailedResponse(http.StatusInternalServerError, "Unable to analyse resume", gin.H{"error": err.Error()}, c)
		return
	}
}
