package leetcode

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/sirupsen/logrus"
)

type LeetClient struct {
	client *graphql.Client
	debug  bool
}

func NewClient(endpoint string) *LeetClient {
	return &LeetClient{
		client: graphql.NewClient(endpoint),
	}
}

func (c *LeetClient) Debug(debug bool) {
	c.debug = debug
	if debug {
		c.client.Log = func(s string) {
			logrus.Debug(s)
		}
	} else {
		c.client.Log = func(s string) {}
	}
}

func (c *LeetClient) GetQuestionOfToday(ctx context.Context) (questions []TodayRecord, err error) {
	request := graphql.NewRequest(`
query questionOfToday {
	todayRecord {
		date
		userStatus
		question {
			questionId
			frontendQuestionId: questionFrontendId
			difficulty
			title
			titleCn: translatedTitle
			titleSlug
			paidOnly: isPaidOnly
			freqBar
			isFavor
			acRate
			status
			solutionNum
			hasVideoSolution
			topicTags {
				name
				nameTranslated: translatedName
				id
			}
			extra {
				topCompanyTags {
				imgUrl
				slug
				numSubscribed
				}
			}
		}
		lastSubmission {
			id
		}
	}
}
	`)

	var response QuestionOfToday
	if err := c.client.Run(ctx, request, &response); err != nil {
		logrus.Errorf("GetQuestionOfToday err: %+v", err)
		return nil, err
	}
	return response.TodayRecord, err
}

func (c *LeetClient) GetRecentSubmissions(ctx context.Context, user string) (submissions []RecentSubmissions, err error) {
	type Data struct {
		RecentSubmissions []RecentSubmissions `json:"recentSubmissions"`
	}

	request := graphql.NewRequest(`
query recentSubmissions($userSlug: String!) {
    recentSubmissions(userSlug: $userSlug) {
        status
        lang
        source {
            sourceType
            ... on SubmissionSrcLeetbookNode {
                slug
                title
                pageId
                __typename
            }
            __typename
        }
        question {
            questionFrontendId
            difficulty
            title
            translatedTitle
            titleSlug
            __typename
        }
        submitTime
        __typename
    }
}
	`)

	request.Var("userSlug", user)

	var response Data
	if err := c.client.Run(ctx, request, &response); err != nil {
		logrus.Errorf("GetRecentSubmissions err: %+v", err)
		return nil, err
	}
	return response.RecentSubmissions, err
}
