package util

import (
	"context"
	"google.golang.org/grpc/metadata"
)

func GetTraceIDFromIncomingContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	val := md.Get("traceID")
	if val == nil || len(val) < 1 {
		return ""
	}
	return val[0]
}
