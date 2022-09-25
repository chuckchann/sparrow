package main

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	order_pb "sparrow/api/protobuf_spec/order"
	config2 "sparrow/config"
	"sparrow/internal/app/order-server-demo/grpc_client"
	"sparrow/internal/app/order-server-demo/rpc_service"
	"sparrow/internal/pkg/middleware/mw_metrics"
	"sparrow/internal/pkg/middleware/mw_rate_limit"
	"sparrow/internal/pkg/middleware/mw_record"
	"sparrow/internal/pkg/middleware/mw_recovery"
	"sparrow/internal/pkg/middleware/mw_trace"
	"sparrow/internal/pkg/monitor"
	"sparrow/internal/pkg/service_mng"
	"sparrow/internal/pkg/slog"
	"sparrow/internal/pkg/util"
)

//调用链路  order_client  --> order_server --> user_server

func main() {

	config2.Init()

	slog.Init(logrus.DebugLevel)

	monitor.Init()

	go monitor.PrometheusServer()

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

	order_pb.RegisterOrderServer(s, rpc_service.NewOrderService())

	//register this server to etcd
	svcMng, err := service_mng.NewServerTypeServiceMng(viper.GetString("namespace"), viper.GetString("appName"))
	if err != nil {
		panic("create service manger err, " + err.Error())
	}

	svcMng.Register(context.Background(), service_mng.NewInstance(util.InternalIp()+":"+viper.GetString("server.port")))

	if err = s.Serve(listen); err != nil {
		panic("serve failed: %s" + err.Error())
	}

}
