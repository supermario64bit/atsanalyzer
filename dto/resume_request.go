package dto

import "mime/multipart"

type ResumeRequest struct {
	JobDescription string                `form:"job_description" binding:"required"`
	ResumeFile     *multipart.FileHeader `form:"resume" binding:"required"`
}

type ResumeAnalysisJob struct {
	JobDescription string `json:"job_description"`
	ResumeTest     string `json:"resume"`
}
