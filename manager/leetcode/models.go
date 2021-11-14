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

type QuestionOfToday struct {
	TodayRecord []TodayRecord `json:"todayRecord"`
}

type TodayRecord struct {
	Date       string `json:"date"`
	UserStatus string `json:"userStatus"`
	Question   struct {
		QuestionID         string      `json:"questionId"`
		FrontendQuestionID string      `json:"frontendQuestionId"`
		Difficulty         string      `json:"difficulty"`
		Title              string      `json:"title"`
		TitleCn            string      `json:"titleCn"`
		TitleSlug          string      `json:"titleSlug"`
		PaidOnly           bool        `json:"paidOnly"`
		FreqBar            interface{} `json:"freqBar"`
		IsFavor            bool        `json:"isFavor"`
		AcRate             float64     `json:"acRate"`
		Status             interface{} `json:"status"`
		SolutionNum        int         `json:"solutionNum"`
		HasVideoSolution   bool        `json:"hasVideoSolution"`
		TopicTags          []struct {
			Name           string `json:"name"`
			NameTranslated string `json:"nameTranslated"`
			ID             string `json:"id"`
		} `json:"topicTags"`
		Extra struct {
			TopCompanyTags []struct {
				ImgURL        string `json:"imgUrl"`
				Slug          string `json:"slug"`
				NumSubscribed int    `json:"numSubscribed"`
			} `json:"topCompanyTags"`
		} `json:"extra"`
	} `json:"question"`
	LastSubmission interface{} `json:"lastSubmission"`
}
