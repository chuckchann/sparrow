package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/spf13/viper"
	"google.golang.org/grpc/resolver"
	"sparrow/internal/pkg/slog"
	"sync"
	"time"
)

const schema = "grpclb"

type Discovery struct {
	cli               *clientv3.Client
	cc                resolver.ClientConn
	instances         sync.Map
	mu                sync.RWMutex
	ttl               int64
	timeout           int64
	needRegisteredSvc []string
	schema            string
}

func NewDiscovery(schema string) resolver.Builder {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("registry.etcd.endpoints"),
		DialTimeout: 1500 * time.Second,
	})
	if err != nil {
		panic(err)
	}

	return &Discovery{
		schema:  schema,
		cli:     cli,
		timeout: 10, //default
	}
}

//Build() put called service info into memory
func (d *Discovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	d.cc = cc
	//url schema:    scheme://authority/endpoint
	prefix := fmt.Sprintf("/%s/%s/", target.Scheme, target.Endpoint)
	resp, err := d.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, v := range resp.Kvs {
		d.setInstance(string(v.Key), string(v.Value))
	}

	//watch service change
	go d.watch(prefix)

	return d, nil
}

func (d *Discovery) Scheme() string {
	return d.schema
}

func (*Discovery) ResolveNow(resolver.ResolveNowOptions) {
	slog.Info("resolve now")
}

func (d *Discovery) Close() {
	d.instances.Range(func(key, value interface{}) bool {
		d.instances.Delete(key)
		return true
	})
	d.cli.Close()
	slog.Info("builder close")
}

func (d *Discovery) getInstanceAddresses() []resolver.Address {
	addrs := make([]resolver.Address, 0)
	d.instances.Range(func(key, value interface{}) bool {
		addrs = append(addrs, value.(resolver.Address))
		return true
	})
	return addrs
}

func (d *Discovery) setInstance(key, value string) {
	d.instances.Store(key, resolver.Address{Addr: value})
	d.cc.UpdateState(resolver.State{Addresses: d.getInstanceAddresses()})
	slog.Debug("builder set instance, key:", key, "value:", value)
}

func (d *Discovery) delInstance(key string) {
	d.instances.Delete(key)
	d.cc.UpdateState(resolver.State{Addresses: d.getInstanceAddresses()})
	slog.Debug("builder delete instance, key:", key)
}

func (d *Discovery) watch(prefix string) {
	watcher := clientv3.NewWatcher(d.cli)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(d.timeout)*time.Second)
	watchChan := watcher.Watch(ctx, prefix, clientv3.WithPrefix())

	//watch the service changes
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			key := string(event.Kv.Key)
			value := string(event.Kv.Value)
			slog.Debugf("ETCDDiscovery: event, type:%s key:%d value:%d", event.Type.String(), key, value)
			switch event.Type {
			case mvccpb.PUT: //add or edit event
				d.setInstance(key, value)
			case mvccpb.DELETE: //delete event
				d.delInstance(string(event.Kv.Key))
			}
		}
	}

	slog.Debug("watch finish ...")
}
