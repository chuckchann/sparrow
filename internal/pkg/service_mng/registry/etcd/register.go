package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/viper"
	"sparrow/internal/pkg/slog"
	"time"
)

type Registry struct {
	//ectd client
	cli *clientv3.Client
	//lease id
	laseID  clientv3.LeaseID
	timeout int64
	ttl     int64
}

func NewRegistry(timeout, ttl int64) (*Registry, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("registry.etcd.endpoints"),
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	r := &Registry{
		cli:     cli,
		timeout: timeout,
		ttl:     ttl,
	}
	return r, nil
}

func (r *Registry) Register(ctx context.Context, namespace, serviceName, ip, port string) {
	kv := clientv3.NewKV(r.cli)
	lease := clientv3.NewLease(r.cli)
	resp, err := lease.Grant(ctx, r.ttl)
	if err != nil {
		slog.Panic("ETCDRegister: grant lase err,", err.Error())
	} else {
		r.laseID = resp.ID
		addr := fmt.Sprintf("%s:%s", ip, port)
		key := fmt.Sprintf("/%s/%s/%s", namespace, serviceName, addr)
		_, err := kv.Put(ctx, key, addr, clientv3.WithLease(r.laseID))
		if err != nil {
			slog.Panic("ETCDRegister: put kv value err,", err)
		}
		ch, err := lease.KeepAlive(ctx, r.laseID)
		if err != nil {
			slog.Panic("ETCDRegister: renewal err, ", err)
		}

		go func() {
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						slog.Errorf("Register: response channel closed, leaseID:%d", r.laseID)
						return
					} else {
						slog.Debugf("Register: lase successfully, leaseID:%d", r.laseID)
					}
				case <-ctx.Done():
					slog.Info("ctx done, ", ctx.Err())
					return
				}
			}
		}()
	}
}

func (r *Registry) Revoke(ctx context.Context) {
	_, err := r.cli.Revoke(ctx, r.laseID)
	if err != nil {
		slog.Warn("Register: revoke lease err,", err.Error())
	}
}
