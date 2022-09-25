package service_mng

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"sparrow/internal/pkg/service_mng/builder"
	"sparrow/internal/pkg/util"
	"time"
)

type Instance struct {
	UID             string           `json:"uid"`
	Address         resolver.Address `json:"endpoint"`
	LatestTimestamp int64            `json:"timestamp"`
	clientConn      *grpc.ClientConn `json:"-"`
	Endpoint        string
}

func (i *Instance) String() string {
	bytes, _ := json.Marshal(i)
	return string(bytes)
}

func NewInstance(addr string) *Instance {
	return &Instance{
		UID:             util.GenUUID(),
		Address:         resolver.Address{Addr: addr},
		LatestTimestamp: time.Now().Unix(),
	}
}

func (i *Instance) GetGRPCConn() *grpc.ClientConn {
	return i.clientConn
}

func genPrefix(namespace, appName string) string {
	return fmt.Sprintf("/%s/%s/%s/", namespace, appName, builder.schema)
}
