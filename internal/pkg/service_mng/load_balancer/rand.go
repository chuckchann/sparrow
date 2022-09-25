package load_balancer

import (
	"context"
	"sparrow/internal/pkg/service_mng"
	"sparrow/internal/pkg/util"
)

type loadBalancer interface {
	Do(ctx context.Context, ins map[string]*service_mng.Instance) *service_mng.Instance
}

type RandomLoadBalance struct {
}

func (*RandomLoadBalance) Do(ctx context.Context, ins map[string]*service_mng.Instance) *service_mng.Instance {
	var insSlice []*service_mng.Instance
	for _, v := range ins {
		insSlice = append(insSlice, v)
	}
	if insSlice == nil {
		return nil
	} else {
		return insSlice[util.RandomIntNum(len(insSlice))]
	}
}
