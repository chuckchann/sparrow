package etcd

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/viper"
	"sparrow/internal/pkg/service_mng"
	"sparrow/internal/pkg/slog"
	"time"
)

type registry interface {
	Name() string
	Do(cxt context.Context, prefix string, ins *service_mng.Instance)
	Revoke(ctx context.Context)
}

type ETCDRegistry struct {
	cli     *clientv3.Client
	laseID  clientv3.LeaseID
	timeout int64
	ttl     int64
}

func NewETCDRegister(timeout, ttl int64) (*ETCDRegistry, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   viper.GetStringSlice("registry.etcd.endpoints"),
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	r := &ETCDRegistry{
		cli:     cli,
		timeout: timeout,
		ttl:     ttl,
	}
	return r, nil
}

func (r *ETCDRegistry) Name() string {
	return "ETCD_REGISTER"
}

func (r *ETCDRegistry) Do(ctx context.Context, prefix string, ins *service_mng.Instance) {
	kv := clientv3.NewKV(r.cli)
	lease := clientv3.NewLease(r.cli)
	resp, err := lease.Grant(ctx, r.ttl)
	if err != nil {
		slog.Panic("ETCDRegister: grant lease err,", err.Error())
	} else {
		//key :=  + ins.Endpoint
		r.laseID = resp.ID
		_, err := kv.Put(ctx, "", ins.String(), clientv3.WithLease(r.laseID))
		if err != nil {
			slog.Warn("ETCDRegister: put kv value err,", err.Error())
			//todo: add error
		}
		ch, err := lease.KeepAlive(ctx, r.laseID)
		if err != nil {
			slog.Warn("ETCDRegister: renewal err, ", err.Error())
			//todo: add error
		}

		go func() {
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						slog.Errorf("ETCDRegister: response channel closed, leaseID:%d", r.laseID)
						return
					}
				case <-ctx.Done():
					slog.Info("sctx done, ", ctx.Err())
					return
				}
			}
		}()
	}
}

func (r *ETCDRegistry) Revoke(ctx context.Context) {
	_, err := r.cli.Revoke(ctx, r.laseID)
	if err != nil {
		slog.Warn("ETCDRegister: revoke lease err,", err.Error())
	}
}
