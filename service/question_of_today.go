package service

import (
	"context"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	"github.com/SsrCoder/leetwatcher/manager/leetcode"
	"github.com/SsrCoder/leetwatcher/manager/opq"
)

var (
	TodayQuestion = leetcode.TodayRecord{}
)

func (s *Service) LoadQuestionOfToday() {
	ctx := context.Background()
	questions, err := s.lc.GetQuestionOfToday(ctx)
	if err != nil {
		logrus.WithContext(ctx).Errorf("GetQuestionOfToday fail: %+v", err)
		return
	}
	if len(questions) != 1 {
		logrus.WithContext(ctx).Errorf("GetQuestionOfToday length != 1, questions: %+v", questions)
		return
	}
	question := questions[0]
	if question.Date != time.Now().Format("2006-01-02") {
		logrus.WithContext(ctx).Errorf("GetQuestionOfToday question is not today: %+v", question.Date)
		return
	}
	TodayQuestion = question
}

func (s *Service) NotifyQuestionOfToday() {
	ctx := context.Background()

	if TodayQuestion.Date != time.Now().Format("2006-01-02") {
		logrus.WithContext(ctx).Errorf("GetQuestionOfToday question is not today: %+v", TodayQuestion.Date)

		// retry once
		s.LoadQuestionOfToday()
		if TodayQuestion.Date != time.Now().Format("2006-01-02") {
			logrus.WithContext(ctx).Errorf("GetQuestionOfToday question is not today: %+v", TodayQuestion.Date)
			return
		}
	}

	s.SendNotifyQuestionOfTodayMessage(ctx)
}

func (s *Service) SendNotifyQuestionOfTodayMessage(ctx context.Context) {
	groups, err := s.bot.FindGroupsByGroupName("LeetWatcher")
	if err != nil {
		logrus.Errorf("FindGroupsByGroupName error: %+v", err)
		return
	}

	msg := s.buildNotifyQuestionOfTodayMessage()

	for _, group := range groups {
		s.bot.SendGroupTextMsg(group.GroupID, msg)
	}
}

func (s *Service) buildNotifyQuestionOfTodayMessage() string {
	builder := strings.Builder{}

	builder.WriteString("每日一题来啦~")
	builder.WriteString("\n")
	builder.WriteString("\n")

	builder.WriteString("题目：")
	builder.WriteString(TodayQuestion.Question.TitleCn)
	builder.WriteString("\n")

	builder.WriteString("难度：")
	builder.WriteString(leetcode.DifficultyMap[TodayQuestion.Question.Difficulty])
	builder.WriteString("\n")

	builder.WriteString("AC率：")
	builder.WriteString(cast.ToString(TodayQuestion.Question.AcRate * 100)[:4] + "%")
	builder.WriteString("\n")

	builder.WriteString("标签：")
	for idx, tag := range TodayQuestion.Question.TopicTags {
		if idx != 0 {
			builder.WriteString("，")
		}
		builder.WriteString(tag.NameTranslated)
	}
	builder.WriteString("\n")

	builder.WriteString("出题公司：")
	for idx, company := range TodayQuestion.Question.Extra.TopCompanyTags {
		if idx != 0 {
			builder.WriteString("，")
		}
		if slug, ok := leetcode.CompanySlugMap[company.Slug]; ok {
			builder.WriteString(slug)
		} else {
			builder.WriteString(company.Slug)
		}
		builder.WriteString("[")
		builder.WriteString(cast.ToString(company.NumSubscribed))
		builder.WriteString("]")
	}
	builder.WriteString("\n")

	builder.WriteString("链接：")
	builder.WriteString("https://leetcode-cn.com/problems/")
	builder.WriteString(TodayQuestion.Question.TitleSlug)

	return builder.String()
}

func (s *Service) NotifyQuestionOfTodayNight() {
	ctx := context.Background()

	if TodayQuestion.Date != time.Now().Format("2006-01-02") {
		logrus.WithContext(ctx).Errorf("NotifyQuestionOfTodayNight question is not today: %+v", TodayQuestion.Date)

		// retry once
		s.LoadQuestionOfToday()
		if TodayQuestion.Date != time.Now().Format("2006-01-02") {
			logrus.WithContext(ctx).Errorf("NotifyQuestionOfTodayNight question is not today: %+v", TodayQuestion.Date)
			return
		}
	}

	s.SendNotifyQuestionOfTodayNightMessage(ctx)
}

func (s *Service) SendNotifyQuestionOfTodayNightMessage(ctx context.Context) {
	groups, err := s.bot.FindGroupsByGroupName("LeetWatcher")
	if err != nil {
		logrus.Errorf("FindGroupsByGroupName error: %+v", err)
		return
	}

	msg := s.buildNotifyQuestionOfTodayNightMessage()

	for _, group := range groups {
		s.bot.SendGroupTextMsg(group.GroupID, msg)
	}
}

func (s *Service) buildNotifyQuestionOfTodayNightMessage() string {
	builder := strings.Builder{}

	builder.WriteString(opq.AtAll())
	builder.WriteString("\n")

	builder.WriteString("都几点了，还不快来刷题？")
	builder.WriteString("\n")
	builder.WriteString("\n")

	builder.WriteString(TodayQuestion.Question.TitleCn)
	builder.WriteString("：")
	builder.WriteString("https://leetcode-cn.com/problems/")
	builder.WriteString(TodayQuestion.Question.TitleSlug)

	return builder.String()
}
