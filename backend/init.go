package backend

import (
	"github.com/SsrCoder/leetwatcher/service"
	"github.com/robfig/cron/v3"
)

type Backend struct {
	s *service.Service
	c *cron.Cron
}

func New(s *service.Service) *Backend {
	return &Backend{
		s: s,
		c: cron.New(cron.WithSeconds()),
	}
}

func (b *Backend) Init() {
	b.s.ReloadUsers()
	b.c.AddFunc("*/15 * * * * *", b.s.ReloadUsersWithRefresh)

	b.s.LoadQuestionOfToday()
	b.c.AddFunc("0 0 * * * *", b.s.LoadQuestionOfToday)
	b.c.AddFunc("0 0 9 * * *", b.s.NotifyQuestionOfToday)       // 早9点提醒
	b.c.AddFunc("0 0 22 * * *", b.s.NotifyQuestionOfTodayNight) // 晚10点提醒
}

func (b *Backend) Start() {
	b.c.Start()
}

func (b *Backend) Stop() {
	b.c.Stop()
}
