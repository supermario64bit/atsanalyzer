package utils

import "google.golang.org/genai"

func GenerateAnalysisSchema() *genai.Schema {
	return &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"ats_match_score": {
				Type:        genai.TypeInteger,
				Description: "ATS match percentage (0-100)",
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
		},
		Required: []string{"ats_match_score", "matched_skills", "missing_skills", "suggestions_to_improve"},
	}
}
