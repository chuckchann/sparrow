package mw_header

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sparrow/internal/pkg/slog"
	"sparrow/internal/pkg/util"
)

//set traceID
func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		var traceID string
		if md.Get("traceID") == nil {
			//generate traceID if there is no traceID in metadata
			traceID = viper.GetString("env") + "-" + util.GenUUID()
			md.Set("traceID", traceID)
			slog.Warn("request metadata do not contains traceID, generate traceID:", traceID)
		} else {
			traceID = md.Get("traceID")[0]
			slog.Debug("get traceID from metadata: ", traceID)
		}

		//set traceID in ctx
		newCtx := context.WithValue(metadata.NewIncomingContext(ctx, md), "traceID", traceID)
		return handler(newCtx, req)
	}
}
