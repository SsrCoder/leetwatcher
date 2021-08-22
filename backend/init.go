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
}

func (b *Backend) Start() {
	b.c.Start()
}

func (b *Backend) Stop() {
	b.c.Stop()
}
