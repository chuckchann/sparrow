package grpc_client

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	user_pb "sparrow/api/protobuf_spec/user"
	"sparrow/internal/pkg/service_mng"
	"sparrow/internal/pkg/slog"
	"strings"
)

//cliMap is contains all grpc client you need in this app
var svcClientMap = map[string]*service_mng.ServiceManager{
	"order-server-demo.default": nil,
}

func parse(key string) (string, string) {
	strs := strings.Split(key, ".")
	if len(strs) < 2 {
		panic("")
	}
	return strs[0], strs[1]
}

func Init() {
	clients := viper.GetStringSlice("grpc.clients")
	fmt.Println("clients----------", clients)
	for _, client := range clients {
		host, ns := parse(client)
		m, err := service_mng.NewClientServiceMng(ns, host)
		if err != nil {
			panic("")
		}
		svcClientMap[client] = m
		slog.Infof("discover %s successfully", client)
		m.DiscoveryAndWatch(context.Background())
	}
}

func GetGRPCClient(ctx context.Context, name string) (interface{}, error) {
	clientSvcMng, ok := svcClientMap[name]
	if !ok {
		return nil, errors.New("unregistered client service, check your service name again")
	}
	ins := clientSvcMng.GetInstance(ctx)
	if ins == nil {
		return nil, errors.New("can not get instance")
	}

	//add client service here...
	switch name {
	case "user-server-demo.default":
		return user_pb.NewUserClient(ins.GetGRPCConn()), nil

	default:
		return nil, errors.New("invalid name")

	}

}
