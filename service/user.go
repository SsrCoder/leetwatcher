package service

import (
	"context"
	"fmt"
	"time"

	"github.com/SsrCoder/leetwatcher/manager/leetcode"
	"github.com/SsrCoder/leetwatcher/model"
	"github.com/SsrCoder/leetwatcher/utils"
	"github.com/sirupsen/logrus"
)

var (
	UserNameMap = make(map[string]model.User)
)

func (s *Service) ReloadUsers() {
	ctx := context.Background()
	users, err := s.d.GetAllUsers(ctx)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GetAllUsers fail: %+v", err)
		return
	}
	tempUserNameMap := make(map[string]model.User)
	for _, user := range users {
		tempUserNameMap[user.Username] = user
	}
	UserNameMap = tempUserNameMap
}

func (s *Service) ReloadUsersWithRefresh() {
	s.ReloadUsers()
	ctx := context.Background()
	for username, user := range UserNameMap {
		time.Sleep(100 * time.Millisecond)
		submissions, err := s.lc.GetRecentSubmissions(ctx, username)
		if err != nil {
			logrus.WithContext(ctx).Errorf("GetRecentSubmissions fail: %+v", err)
			continue
		}
		if len(submissions) == 0 {
			continue
		}
		if user.LastSubmitTime.Unix() == submissions[0].SubmitTime {
			continue
		}
		// 入库 & 发消息
		user.LastSubmitTime = time.Unix(submissions[0].SubmitTime, 0)
		if err := s.d.UpdateUserLastSubmitTime(ctx, user.ID, submissions[0].SubmitTime); err != nil {
			logrus.WithContext(ctx).Errorf("UpdateUserLastSubmitTime: %+v", err)
		} else {
			s.SendLeetCodeSubmitMessage(user, submissions[0])
		}
	}
}

func (s *Service) SendLeetCodeSubmitMessage(user model.User, submission leetcode.RecentSubmissions) {
	groupList, err := s.bot.GetGroupList("")
	for len(groupList.TroopList) != 0 {
		if err != nil {
			// TODO: error check
			return
		}

		for _, group := range groupList.TroopList {
			if group.GroupName == "LeetWatcher" {
				if len(leetcode.SubmitStatusMap[submission.Status]) == 0 {
					logrus.WithField("submission", submission).Errorf("Submission status not found | status: %+v", submission.Status)
				}
				if len(submission.Status) == 0 {
					continue
				}
				s.bot.SendGroupTextMsg(group.GroupID, fmt.Sprintf("%s上次提交的题目是：\n\nID: %v\n标题: %s\n状态：%s\n语言：%s\n时间: %s",
					user.Remark, submission.Question.QuestionFrontendID, submission.Question.TranslatedTitle,
					leetcode.SubmitStatusMap[submission.Status], leetcode.SubmitLanguageMap[submission.Lang],
					utils.DatetimeFormat(user.LastSubmitTime)))
			}
		}

		if len(groupList.NextToken) == 0 {
			break
		}
		groupList, err = s.bot.GetGroupList(groupList.NextToken)
	}
}
