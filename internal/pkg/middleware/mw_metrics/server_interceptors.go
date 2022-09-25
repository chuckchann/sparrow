package mw_metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	common_pb "sparrow/api/protobuf_spec/common"
	"sparrow/internal/pkg/monitor"
	"sparrow/internal/pkg/slog"
	"strconv"
	"time"
)

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !viper.GetBool("prometheus.isOpen") {
			return handler(ctx, req)
		}

		start := time.Now()
		resp, err := handler(ctx, req)
		r, flag := resp.(*common_pb.Response)
		if !flag {
			slog.Warn("response is illegal")
		}

		monitor.GetRequestMetrics().With(prometheus.Labels{
			"namespace": viper.GetString("namespace"),
			"app":       viper.GetString("appName"),
			"protocol":  "rpc_service",
			"method":    info.FullMethod,
			"resp_code": strconv.FormatInt(r.Code, 10),
		}).Inc()

		monitor.GetLatencyMetrics().With(prometheus.Labels{
			"namespace": viper.GetString("namespace"),
			"app":       viper.GetString("appName"),
			"protocol":  "rpc_service",
			"method":    info.FullMethod,
			"resp_code": strconv.FormatInt(r.Code, 10),
		}).Observe(float64(time.Now().Sub(start).Milliseconds()))

		return resp, err
	}
}
