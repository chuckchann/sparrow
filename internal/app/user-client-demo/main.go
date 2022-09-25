package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/resolver"
	order_pb "sparrow/api/protobuf_spec/order"
	user_pb "sparrow/api/protobuf_spec/user"
	"sparrow/config"
	"sparrow/internal/pkg/service_mng/discovery/etcd"
	"sparrow/internal/pkg/slog"
	"time"
)

func main() {

	config.Init()
	slog.Init(logrus.DebugLevel)

	b := etcd.NewDiscovery( "dns")
	resolver.Register(b)
	//建立长连接
	conn, err := grpc.Dial(b.Scheme()+"://8.8.8.8/simple_grp", grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		fmt.Println("dial", err)
		return
	}

	//simple way
	callUser(conn)

	//callUser(conn)

	defer conn.Close()

}

func callOrder(conn *grpc.ClientConn) {

	ctx := metadata.AppendToOutgoingContext(context.Background(), "hello", "world")

	c := order_pb.NewOrderClient(conn)
	req := &order_pb.GetOrderInfoReq{OrderID: "xxxx"}

	resp, err := c.GetOrderInfo(ctx, req)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
		return
	}
	order := &order_pb.OrderInfo{}
	fmt.Println(resp.Code, resp.Msg, resp.Data)
	if resp.Code == 0 {
		err = ptypes.UnmarshalAny(resp.Data, order)
		if err != nil {
			fmt.Println("UnmarshalAny 失败", err.Error())
			return
		}
	}

	fmt.Println(resp.Code, resp.Msg, order)
}

func callUser(conn *grpc.ClientConn) {

	c_, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := metadata.AppendToOutgoingContext(c_, "lover", "yejie")

	c := user_pb.NewUserClient(conn)
	req := &user_pb.GetUserInfoRequest{Id: 1}

	resp, err := c.GetUserInfo(ctx, req)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
		return
	}
	user := &user_pb.UserInfo{}
	fmt.Println(resp.Code, resp.Msg, resp.Data)
	if resp.Code == 0 {
		err = ptypes.UnmarshalAny(resp.Data, user)
		if err != nil {
			fmt.Println("UnmarshalAny 失败", err.Error())
			return
		}
	}

	fmt.Println(resp.Code, resp.Msg, user)

}
