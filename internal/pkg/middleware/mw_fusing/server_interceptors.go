package mw_fusing

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"sparrow/internal/pkg/slog"
)

//grpc_server interceptor
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if !viper.GetBool("fusing.isOpen") {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		hystrix.Do(viper.GetString("appName"), func() error {
			return invoker(ctx, method, req, reply, cc, opts...)
		}, func(err error) error {
			slog.Warn("hystrix do err, ", err.Error())
			//TODO: report prometheus
			return nil
		})
		return nil
	}

}
