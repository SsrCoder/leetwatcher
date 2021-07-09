package leetcode

type Question struct {
	QuestionFrontendID string `json:"questionFrontendId"`
	Title              string `json:"title"`
	TranslatedTitle    string `json:"translatedTitle"`
	TitleSlug          string `json:"titleSlug"`
	Typename           string `json:"__typename"`
}

type RecentSubmissions struct {
	Status     string      `json:"status"`
	Lang       string      `json:"lang"`
	Source     interface{} `json:"source"`
	Question   Question    `json:"question"`
	SubmitTime int64       `json:"submitTime"`
	Typename   string      `json:"__typename"`
}
