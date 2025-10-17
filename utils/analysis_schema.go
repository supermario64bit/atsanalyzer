package utils

import "google.golang.org/genai"

func GenerateAnalysisSchema() *genai.Schema {
	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"ats_match_score": {
				Type:        genai.TypeInteger,
				Description: "ATS match percentage (0-100). Match criteria will be the required skills, experience and background in the job description matches with the candidate resume",
			},
			"matched_skills": {
				Type:        genai.TypeArray,
				Description: "A list of skills found in the resume that match the job description.",
				Items:       &genai.Schema{Type: genai.TypeString},
			},
			"missing_skills": {
				Type:        genai.TypeArray,
				Description: "A list of skills from the job description that are missing from the resume.",
				Items:       &genai.Schema{Type: genai.TypeString},
			},
			"suggestions_to_improve": {
				Type:        genai.TypeArray,
				Description: "Specific suggestions to enhance the resume for a better match with the job description in simple human readable english.",
				Items:       &genai.Schema{Type: genai.TypeString},
			},
			"candidate_name": {
				Type:        genai.TypeString,
				Description: "Candidate name given in the resume. If name not available show, 'Unable to extract candidate name from resume'",
			},
			"candidate_email": {
				Type:        genai.TypeString,
				Description: "Candidate email given in the resume. If email not available show, 'Unable to extract candidate email from resume.'",
			},
			"candidate_phone": {
				Type:        genai.TypeString,
				Description: "Candidate phone number given in the resume. If number not available show, 'Unable to extract candidate phone from resume.'",
			},
			"company_name": {
				Type:        genai.TypeString,
				Description: "Company name given in the job description. If not available show, 'Unable to extract company name from JD.'",
			},
			"role": {
				Type:        genai.TypeString,
				Description: "Job Role given in the job description that the company is hiring. If not available show, 'Unable to extract job role from JD.'",
			},
		},
		Required: []string{"ats_match_score", "matched_skills", "missing_skills", "suggestions_to_improve", "candidate_name", "candidate_email", "candidate_phone", "company_name", "role"},
	}
}
