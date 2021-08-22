package main

import (
	"flag"

	"github.com/SsrCoder/leetwatcher/backend"
	"github.com/SsrCoder/leetwatcher/config"
	"github.com/SsrCoder/leetwatcher/dao"
	"github.com/SsrCoder/leetwatcher/manager/opq"
	"github.com/SsrCoder/leetwatcher/service"
	"github.com/SsrCoder/leetwatcher/utils"
)

func init() {
	utils.InitLogrus()
}

var (
	configPath = flag.String("config", "./conf/config.toml", "Configuration")
)

func main() {
	flag.Parse()
	conf, err := config.New(*configPath)
	if err != nil {
		panic(err)
	}
	d, err := dao.New(conf.DB)
	if err != nil {
		panic(err)
	}
	bot := opq.NewBot(1721538083, "http://ssrcoder.com:23300")
	s := service.New(d, bot)
	s.InitBot()
	b := backend.New(s)
	b.Init()
	b.Start()
	defer b.Stop()

	bot.Wait()
}
