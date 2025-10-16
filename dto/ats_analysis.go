package dto

type ResumeAnalysis struct {
    AtsScore            int      `json:"ats_match_score"`
    MatchingSkills      []string `json:"matched_skills"`
    MissingSkills       []string `json:"missing_skills"`
    SuggestionsToImprove []string `json:"suggestions_to_improve"`
}
