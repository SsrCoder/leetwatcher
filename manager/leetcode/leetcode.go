package leetcode

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

type leetClient struct {
	client *graphql.Client
	debug  bool
}

func NewClient(endpoint string) *leetClient {
	return &leetClient{
		client: graphql.NewClient(endpoint),
	}
}

func (c *leetClient) Debug(debug bool) {
	c.debug = debug
	if debug {
		c.client.Log = func(s string) {
			fmt.Println(s)
		}
	} else {
		c.client.Log = func(s string) {}
	}
}

func (c *leetClient) GetRecentSubmissions(ctx context.Context, user string) (submissions []RecentSubmissions, err error) {
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
		return nil, err
	}
	return response.RecentSubmissions, err
}
