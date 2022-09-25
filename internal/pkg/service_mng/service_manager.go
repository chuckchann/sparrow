package service_mng

import (
	"context"
	"fmt"
	discovery2 "sparrow/internal/pkg/service_mng/discovery"
	registry2 "sparrow/internal/pkg/service_mng/registry"
	"sparrow/internal/pkg/service_mng/registry/etcd"
)

const (
	TYPE_SERVER = iota
	TYPE_CLIENT
)

type SvcMngType int

func (smt SvcMngType) String() string {
	if smt == TYPE_SERVER {
		return "grpc_server type service manger"
	} else if smt == TYPE_CLIENT {
		return "grpc_client type service manger"
	}
	return ""
}

type ServiceManager struct {
	svcMngType SvcMngType
	Namespace  string
	Host       string
	registry2.registry
	discovery2.discovery
}

func NewServerTypeServiceMng(namespace, host string) (*ServiceManager, error) {
	r, err := etcd.NewETCDRegister(5, 10)
	if err != nil {
		return nil, err
	}
	return &ServiceManager{
		svcMngType: TYPE_SERVER, //server type manger
		Namespace:  namespace,
		Host:       host,
		registry:   r,
	}, nil
}

func NewClientServiceMng(namespace, host string) (*ServiceManager, error) {
	d, err := discovery2.NewETCDDiscovery(10)
	if err != nil {
		return nil, err
	}
	return &ServiceManager{
		svcMngType: TYPE_CLIENT, //server type manger
		Namespace:  namespace,
		Host:       host,
		discovery:  d,
	}, nil
}

func (sm *ServiceManager) Type() SvcMngType {
	return sm.svcMngType
}

func (sm *ServiceManager) genPrefix() string {
	return fmt.Sprintf("/%s/%s/", sm.Namespace, sm.Host)
}

func (sm *ServiceManager) Register(ctx context.Context, ins *Instance) {
	sm.registry.Do(ctx, sm.genPrefix(), ins)
}

func (sm *ServiceManager) DiscoveryAndWatch(ctx context.Context) {
	sm.discovery.Do(ctx, sm.genPrefix())
}

func (sm *ServiceManager) GetInstance(ctx context.Context) *Instance {
	return sm.discovery.GetInstance(ctx)
}
