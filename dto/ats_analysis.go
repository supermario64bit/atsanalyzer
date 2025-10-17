package dto

type ResumeAnalysis struct {
	IsResume             bool     `json:"is_resume"`
	IsJD                 bool     `json:"is_jd"`
	AtsScore             int      `json:"ats_match_score"`
	MatchingSkills       []string `json:"matched_skills"`
	MissingSkills        []string `json:"missing_skills"`
	SuggestionsToImprove []string `json:"suggestions_to_improve"`
	CandiateName         string   `json:"candidate_name"`
	CandidateEmail       string   `json:"candidate_email"`
	CandidatePhone       string   `json:"candidate_phone"`
	CompanyName          string   `json:"company_name"`
	JobRole              string   `json:"role"`
}
