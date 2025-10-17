package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/utils"
	"google.golang.org/genai"
)

func AnalyzeResumeWithJD(req *dto.ResumeRequest) (*dto.ResumeAnalysis, error) {
	resumeText, err := utils.ExtractTextFromPDF(req.ResumeFile)
	if err != nil {
		return nil, err
	}

	genAIClient, err := NewClient()
	if err != nil {
		return nil, fmt.Errorf("Unable initialize ai client. Error: " + err.Error())
	}

	prompt := fmt.Sprintf(`
                Analyze the provided Resume against the Job Description.
                Job Description: %s
                Resume: %s
                Provide a detailed analysis following the exact JSON schema.
                `, req.JobDescription, resumeText)

	parts := []*genai.Part{
		{Text: prompt},
	}

	temp := float32(0.0)
	topP := float32(0.9)
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema:   utils.GenerateAnalysisSchema(),
		Temperature:      &temp,
		TopP:             &topP,
	}
	resp, err := genAIClient.Models.GenerateContent(context.Background(), "gemini-2.5-flash", []*genai.Content{{Role: "user", Parts: parts}}, config)
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("received empty or invalid content response from AI model")
	}

	jsonString := resp.Candidates[0].Content.Parts[0].Text

	var analysis dto.ResumeAnalysis
	err = json.Unmarshal([]byte(jsonString), &analysis)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response into struct: %w. Response text: %s", err, jsonString)
	}

	return &analysis, err

}
