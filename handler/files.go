package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/services"
)

type FileHandler interface {
	UploadFile(c *gin.Context)
}

type fileHandler struct{}

func NewFileHandler() FileHandler {
	return &fileHandler{}
}

func (h *fileHandler) UploadFile(c *gin.Context) {
	var dto dto.ResumeRequest

	if err := c.ShouldBind(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid request body", "result": gin.H{"error": err.Error()}})
		return
	}

	cloudinarySvc, err := services.NewCloudinaryService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "Cloudinary Intialization Failed", "result": gin.H{"error": err.Error()}})
		return
	}

	url, publicId, err := cloudinarySvc.UploadPdf(dto.ResumeFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": "File Upload Failed", "result": gin.H{"error": err.Error()}})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "success", "message": "File Uploaded!!", "result": gin.H{"url": url, "public_id": publicId}})
}
