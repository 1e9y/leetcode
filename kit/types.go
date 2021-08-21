package kit

type AllData struct {
	Username        string                   `json:"user_name"`
	StatStatusPairs []AllDataStatStatusPairs `json:"stat_status_pairs"`
}

type AllDataStatStatusPairs struct {
	Stat       AllDataStat       `json:"stat"`
	Difficulty AllDataDifficulty `json:"difficulty"`
}

type AllDataStat struct {
	QuestionID        int    `json:"question_id,omitempty"`
	QuestionTitle     string `json:"question__title,omitempty"`
	QuestionTitleSlug string `json:"question__title_slug,omitempty"`
}

type AllDataDifficulty struct {
	Level int `json:"level"`
}

type QuestionData struct {
	Data QuestionDataData `json:"data"`
}

type QuestionDataData struct {
	Question QuestionDataQuestion `json:"question"`
}

type QuestionDataCodeSnippet struct {
	Lang     string `json:"lang,omitempty"`
	LangSlug string `json:"langSlug,omitempty"`
	Code     string `json:"code,omitempty"`
}

type QuestionDataQuestion struct {
	QuestionID   string                    `json:"questionId"`
	Title        string                    `json:"title"`
	TitleSlug    string                    `json:"titleSlug"`
	Content      string                    `json:"content"`
	CodeSnippets []QuestionDataCodeSnippet `json:"codeSnippets"`
}
