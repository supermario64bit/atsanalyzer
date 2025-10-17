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
		You are an ATS evaluation assistant.

		Your job has two parts:

		**Step 1: Classification**
		Determine independently for each text:
		- If the FIRST text resembles a job post, job listing, or hiring ad — set is_jd=true.
		- If the SECOND text resembles a resume, CV, or candidate profile — set is_resume=true.

		Use these rules to decide:
		- Job Descriptions usually contain words like “Responsibilities”, “Requirements”, “Role”, “We are hiring”, or “Job Title”.
		- Resumes usually contain sections like “Experience”, “Projects”, “Education”, contact details (email/phone), bullet points, or company names.
		- If the text is generic, meaningless, or lorem ipsum, both should be false.

		**Step 2: ATS Analysis**
		Only if BOTH is_jd and is_resume are true:
		Perform an ATS match and fill out the remaining fields.
		If either one is false:
		Return descriptive placeholder values like "JD invalid" or "Resume invalid" for all other fields.

		FIRST TEXT (potential Job Description):
		%s

		SECOND TEXT (potential Resume):
		%s

		Output JSON strictly following this schema:
		{
		"is_resume": boolean,
		"is_jd": boolean,
		"ats_match_score": integer (0-100),
		"matched_skills": [string],
		"missing_skills": [string],
		"suggestions_to_improve": [string],
		"candidate_name": string,
		"candidate_email": string,
		"candidate_phone": string,
		"company_name": string,
		"role": string
		}
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
