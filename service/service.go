package service

import (
	"github.com/mcoo/OPQBot"
	"github.com/sirupsen/logrus"

	"github.com/SsrCoder/leetwatcher/dao"
	"github.com/SsrCoder/leetwatcher/manager/leetcode"
	"github.com/SsrCoder/leetwatcher/manager/opq"
)

type Service struct {
	d   *dao.Dao
	lc  *leetcode.LeetClient
	bot *opq.Bot
}

func New(d *dao.Dao, bot *opq.Bot) *Service {
	client := leetcode.NewClient("https://leetcode-cn.com/graphql/")
	client.Debug(true)
	return &Service{
		d:   d,
		bot: bot,
		lc:  client,
	}
}

func (s *Service) InitBot() {
}

func checkSender(botQQ int64, packet *OPQBot.GroupMsgPack) {
	if packet.FromUserID == botQQ {
		return
	}
	packet.Next(botQQ, packet)
}

func checkGroup(botQQ int64, packet *OPQBot.GroupMsgPack) {
	if packet.FromGroupName != "LeetWatcher" {
		return
	}
	packet.Next(botQQ, packet)
}

func Hello() {
	logrus.Info("start")
	logrus.Debug("start")
	logrus.Warn("start")
	logrus.Error("start")
}
