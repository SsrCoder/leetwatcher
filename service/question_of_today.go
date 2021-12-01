package service

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"

	"github.com/SsrCoder/leetwatcher/manager/leetcode"
	"github.com/SsrCoder/leetwatcher/manager/opq"
	"github.com/SsrCoder/leetwatcher/model"
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
	xmlMsg := s.buildNotifyQuestionOfTodayXMLMessage()

	for _, group := range groups {
		s.bot.SendGroupTextMsg(group.GroupID, msg)
		s.bot.SendGroupXmlMsg(group.GroupID, xmlMsg)
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

func (s *Service) buildNotifyQuestionOfTodayXMLMessage() string {
	xml, err := template.ParseFiles("conf/xml/question_of_today.xml")
	if err != nil {
		panic(err)
	}

	builder := strings.Builder{}
	for idx, tag := range TodayQuestion.Question.TopicTags {
		if idx != 0 {
			builder.WriteString("，")
		}
		builder.WriteString(tag.NameTranslated)
	}
	labels := builder.String()

	builder.Reset()
	for idx, company := range TodayQuestion.Question.Extra.TopCompanyTags {
		if idx != 0 {
			builder.WriteString(",")
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
	companies := builder.String()

	now := time.Now()
	msgModel := model.QuestionOfTodayMessageModel{
		URL:        "https://leetcode-cn.com/problems/" + TodayQuestion.Question.TitleSlug,
		Title:      TodayQuestion.Question.TitleCn,
		Difficulty: leetcode.DifficultyMap[TodayQuestion.Question.Difficulty],
		AcRate:     cast.ToString(TodayQuestion.Question.AcRate * 100)[:4] + "%",
		Labels:     labels,
		Companies:  companies,
		Picture:    fmt.Sprintf("https://assets.leetcode-cn.com/medals/%v/lg/%v-%v.png", now.Year(), now.Year(), int(now.Month())),
	}

	buf := &bytes.Buffer{}
	xml.Execute(buf, msgModel)
	return buf.String()
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
