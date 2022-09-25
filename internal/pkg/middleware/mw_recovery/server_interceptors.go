package mw_recovery

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"runtime"
	"sparrow/internal/pkg/serror"

	"sparrow/internal/pkg/slog"

	"google.golang.org/grpc"
)

func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if pErr := recover(); err != nil {
				buf := make([]byte, 2<<10)
				runtime.Stack(buf, false)
				slog.Error("process panic:", pErr, string(buf))
				resp = serror.ERR_SYSTEM.ReplaceMsg("server panic").Response()
				err = status.Error(codes.Internal, "server panic")
				return
			}
		}()

		return handler(ctx, req)
	}
}
