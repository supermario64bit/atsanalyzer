package handler

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request body", "result": gin.H{"error": err.Error()}})
		return
	}

	if len(dto.JobDescription) < 20 || dto.ResumeFile == nil {
		fmt.Println("Empty Request")
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request body", "result": gin.H{"error": "Empty request"}})
		return
	}

	rawText, err := service.AnalyzeResumeWithJD(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Unable to parse resume", "result": gin.H{"error": err.Error()}})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "File Uploaded!!", "result": gin.H{"pdf_raw": rawText}})
}
