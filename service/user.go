package service

import (
	"context"
	"strings"
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
		if len(leetcode.SubmitStatusMap[submissions[0].Status]) == 0 {
			logrus.WithField("submission", submissions[0]).Errorf("Submission status not found | status: %+v", submissions[0].Status)
		}
		if len(submissions[0].Status) == 0 {
			return
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
	groups, err := s.bot.FindGroupsByGroupName("LeetWatcher")
	if err != nil {
		logrus.Errorf("FindGroupsByGroupName error: %+v", err)
		return
	}

	msg := s.buildLeetCodeSubmitMessage(user, submission)

	for _, group := range groups {
		s.bot.SendGroupTextMsg(group.GroupID, msg)
	}
}

func (s *Service) buildLeetCodeSubmitMessage(user model.User, submission leetcode.RecentSubmissions) string {
	builder := &strings.Builder{}
	builder.WriteString(user.Remark)
	builder.WriteString("上次提交的题目是：\n\n")

	builder.WriteString("ID：")
	builder.WriteString(submission.Question.QuestionFrontendID)
	builder.WriteString("\n")

	builder.WriteString("标题：")
	builder.WriteString(submission.Question.Title)
	builder.WriteString("\n")

	builder.WriteString("状态：")
	builder.WriteString(leetcode.SubmitStatusMap[submission.Status])
	builder.WriteString("\n")

	builder.WriteString("语言：")
	builder.WriteString(leetcode.SubmitLanguageMap[submission.Lang])
	builder.WriteString("\n")

	builder.WriteString("时间：")
	builder.WriteString(utils.DatetimeFormat(user.LastSubmitTime))
	builder.WriteString("\n")

	return builder.String()
}
