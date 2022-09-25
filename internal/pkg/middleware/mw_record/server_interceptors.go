package mw_record

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sparrow/internal/pkg/slog"
	"sparrow/internal/pkg/util"
	"time"
)

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println("---------------------------------------test--------------------------------")
		var errMsg string
		resp, err := handler(ctx, req)
		if err != nil {
			errMsg = err.Error()
		}

		md, _ := metadata.FromIncomingContext(ctx)
		traceID := util.GetTraceIDFromIncomingContext(ctx)

		start := time.Now()
		version := viper.GetString("appName") + "/" + viper.GetString("version")
		fields := logrus.Fields{
			"namespace":  viper.GetString("namespace"),
			"app":        viper.GetString("appName"),
			"version":    version,
			"X-Trace-ID": traceID,
			"protocol":   "grpc",
			"header":     md,
			"req":        req,
			"method":     info.FullMethod,
			"server":     info.Server,
			"start":      start,
			"end":        time.Now(),
			"cost":       time.Now().Sub(start),
			"handleErr":  errMsg,
			"response":   resp,
		}
		slog.Info(fields)
		return resp, err
	}
}
