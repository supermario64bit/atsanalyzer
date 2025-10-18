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

func GetPromptResponseThroughGoogleAiStudioAPI(prompt string) (*dto.ResumeAnalysis, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

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

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		[]*genai.Content{{Role: "user", Parts: parts}},
		config,
	)

	if len(result.Candidates) == 0 || result.Candidates[0].Content == nil || len(result.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("received empty or invalid content response from AI model")
	}

	jsonString := result.Candidates[0].Content.Parts[0].Text

	var analysis dto.ResumeAnalysis
	err = json.Unmarshal([]byte(jsonString), &analysis)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response into struct: %w. Response text: %s", err, jsonString)
	}

	return &analysis, err
}
