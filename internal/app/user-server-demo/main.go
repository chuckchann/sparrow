package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	user_pb "sparrow/api/protobuf_spec/user"
	"sparrow/config"
	"sparrow/internal/app/order-server-demo/grpc_client"
	"sparrow/internal/app/user-server-demo/service_biz/interfaces/grpc_handler"
	"sparrow/internal/pkg/fuse"
	"sparrow/internal/pkg/middleware/mw_metrics"
	"sparrow/internal/pkg/middleware/mw_rate_limit"
	"sparrow/internal/pkg/middleware/mw_record"
	"sparrow/internal/pkg/middleware/mw_recovery"
	"sparrow/internal/pkg/middleware/mw_trace"
	"sparrow/internal/pkg/monitor"
	"sparrow/internal/pkg/service_mng/registry/etcd"
	"sparrow/internal/pkg/slog"
	"sparrow/internal/pkg/util"
)

func main() {
	//init base_component
	config.Init()

	slog.Init(logrus.InfoLevel)

	//init prometheus monitor
	monitor.Init()

	//prometheus server
	go monitor.PrometheusServer()

	//pprof server
	go monitor.PProfServer()

	//fuse init
	fuse.Init()

	grpc_client.Init()

	run()
}

func run() {
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if host == "" || port == "" {
		panic("address invalid")
	}
	addr := host + ":" + port
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			//grpc_recovery.UnaryServerInterceptor(),
			mw_recovery.UnaryInterceptor(),         //panic recovery
			mw_record.UnaryInterceptor(),           //request record,
			mw_rate_limit.UnaryServerInterceptor(), //rate limit
			mw_metrics.UnaryInterceptor(),          //prometheus metrics
			mw_trace.UnaryServerInterceptor(),      //trace
		)),
	)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("listen failed: %s", err.Error()))
	}

	user_pb.RegisterUserServer(s, grpc_handler.NewUserServer(nil))

	r, err := etcd.NewRegistry(30, 10)
	if err != nil {
		slog.Panic(err)
	}
	//resister service message
	r.Register(
		context.Background(),
		viper.GetString("namespace"),
		viper.GetString("appName"),
		util.InternalIp(),
		viper.GetString("server.port"))

	if err = s.Serve(listen); err != nil {
		panic(fmt.Sprintf("serve failed: %s", err.Error()))
	}

}
