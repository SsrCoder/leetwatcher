package opq

import (
	"github.com/mcoo/OPQBot"
)

type bot struct {
	bm *OPQBot.BotManager
}

func NewBot(qq int64, host string) *bot {
	bm := OPQBot.NewBotManager(qq, host)
	return &bot{
		bm: &bm,
	}
}

func (b *bot) OnCrontab(crontab string, fn func()) {

}

func (b *bot) OnGroupMessage(funcs ...GroupMessageFunc) {
	var f []interface{}
	for _, ff := range funcs {
		f = append(f, ff)
	}
	b.bm.AddEvent(OPQBot.EventNameOnGroupMessage, f...)
}

func (b *bot) Wait() {
	b.bm.Wait()
}
