package main

import (
	"flag"
	"os"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/conf"

	"github.com/bluven/f-cloud/app/instance/taskq/internal/config"
	"github.com/bluven/f-cloud/app/instance/taskq/internal/handler"
	"github.com/bluven/f-cloud/app/instance/taskq/internal/svc"
)

var configFile = flag.String("f", "etc/taskq.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	mux := handler.Register(svcContext)

	server := newAsyncServer(c)
	if err := server.Run(mux); err != nil {
		os.Exit(1)
	}
}

func newAsyncServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass},
		asynq.Config{
			Concurrency: 10,
		},
	)
}
