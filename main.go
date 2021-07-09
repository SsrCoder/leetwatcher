package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SsrCoder/leetwatcher/manager/leetcode"
	"github.com/SsrCoder/leetwatcher/manager/opq"
	"github.com/mcoo/OPQBot"
)

func main() {
	bot := opq.NewBot(1721538083, "http://ssrcoder.com:23300")
	bot.OnGroupMessage(checkGroup, checkSender, handleGroupMessage)
	bot.Wait()
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

func handleGroupMessage(botQQ int64, packet *OPQBot.GroupMsgPack) {
	switch packet.Content {
	case "show":
		client := leetcode.NewClient("https://leetcode-cn.com/graphql/")
		client.Debug(false)
		submissions, err := client.GetRecentSubmissions(context.Background(), "feng-zhong-zhui-feng")
		if err != nil {
			panic(err)
		}
		if len(submissions) > 0 {
			submitTime := time.Unix(submissions[0].SubmitTime, 0)
			packet.Bot.SendGroupTextMsg(packet.FromGroupID, fmt.Sprintf("裴神上次提交的题目是：\n\nID: %v\n标题: %s\n时间: %s", submissions[0].Question.QuestionFrontendID, submissions[0].Question.TranslatedTitle, DatetimeFormat(submitTime)))
		} else {
			packet.Bot.SendGroupTextMsg(packet.FromGroupID, "No submissions")
		}
	default:
		packet.Bot.SendGroupTextMsg(packet.FromGroupID, packet.Content)
	}
}

func DatetimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
