package main

import (
	"flag"
	"fmt"

	"github.com/bluven/f-cloud/app/network/rpc/internal/config"
	"github.com/bluven/f-cloud/app/network/rpc/internal/server"
	"github.com/bluven/f-cloud/app/network/rpc/internal/svc"
	"github.com/bluven/f-cloud/app/network/rpc/proto"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/network.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		proto.RegisterNetworkServer(grpcServer, server.NewNetworkServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
