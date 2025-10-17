package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/service"
)

type FileHandler interface {
	Analyse(c *gin.Context)
}

type fileHandler struct{}

func NewFileHandler() FileHandler {
	return &fileHandler{}
}

func (h *fileHandler) Analyse(c *gin.Context) {
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

	rawText, err := service.AnalyzeResumeWithJD(&dto)
	if err != nil {
		log.Println("Unable to parse resume. Error: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Unable to parse resume", "result": gin.H{"error": err.Error()}})
		return
	}

	log.Println("ATS Score available")
	c.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "Analysis completed!!", "result": gin.H{"pdf_raw": rawText}})
}
