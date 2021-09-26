package opq

import (
	"github.com/mcoo/OPQBot"
	"github.com/robfig/cron/v3"
)

type Bot struct {
	*OPQBot.BotManager
	bm *OPQBot.BotManager
	c  *cron.Cron
}

func NewBot(qq int64, host string) *Bot {
	bm := OPQBot.NewBotManager(qq, host)
	return &Bot{
		BotManager: &bm,
		bm:         &bm,
		c:          cron.New(cron.WithSeconds()),
	}
}

func (b *Bot) OnCrontab(crontab string, fn func()) {
	b.c.AddFunc(crontab, fn)
}

func (b *Bot) OnGroupMessage(funcs ...GroupMessageFunc) {
	var f []interface{}
	for _, ff := range funcs {
		f = append(f, ff)
	}
	b.bm.AddEvent(OPQBot.EventNameOnGroupMessage, f...)
}

func (b *Bot) Wait() {
	b.bm.Wait()
}

func (b *Bot) FindGroupsByGroupName(groupName string) (groups []Group, err error) {
	groupList, err := b.GetGroupList("")
	for len(groupList.TroopList) != 0 {
		if err != nil {
			return nil, err
		}

		for _, group := range groupList.TroopList {
			if group.GroupName == groupName {
				groups = append(groups, group)
			}
		}

		if len(groupList.NextToken) == 0 {
			break
		}
		groupList, err = b.GetGroupList(groupList.NextToken)
	}

	return
}
