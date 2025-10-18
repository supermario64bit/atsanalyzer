package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/supermario64bit/atsanalyzer/dto"
	"github.com/supermario64bit/atsanalyzer/types"
)

type resumeAnalysisService struct {
	genAiSvc GenAIService
}

type ResumeAnalysis interface {
	Analyse(req *dto.ResumeAnalysisJob) (*dto.ResumeAnalysis, *types.ApplicationError)
}

func NewResumeAnalysisService() ResumeAnalysis {
	return &resumeAnalysisService{
		genAiSvc: NewGenAiService(),
	}
}

func (svc *resumeAnalysisService) Analyse(req *dto.ResumeAnalysisJob) (*dto.ResumeAnalysis, *types.ApplicationError) {
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
		`, req.JobDescription, req.ResumeTest)

	result, appErr := svc.genAiSvc.GetPromptResponse(prompt)
	if appErr != nil {
		return nil, appErr
	}
	fmt.Errorf("received empty or invalid content response from AI model")

	if len(result.Candidates) == 0 || result.Candidates[0].Content == nil || len(result.Candidates[0].Content.Parts) == 0 {
		return nil, &types.ApplicationError{
			HttpStatusCode: http.StatusInternalServerError,
			Message:        "Unable to analyse resume",
			Error:          fmt.Errorf("received empty or invalid content response from AI model"),
		}
	}

	jsonString := result.Candidates[0].Content.Parts[0].Text

	var analysis dto.ResumeAnalysis
	err := json.Unmarshal([]byte(jsonString), &analysis)
	if err != nil {
		return nil, &types.ApplicationError{
			HttpStatusCode: http.StatusInternalServerError,
			Message:        "Unable to analyse resume",
			Error:          fmt.Errorf("failed to unmarshal JSON response into struct: %w. Response text: %s", err, jsonString),
		}
	}

	return &analysis, nil
}
