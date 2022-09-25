package mw_trace

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sparrow/api/protobuf_spec/common"
	"sparrow/internal/pkg/slog"
)

type TextMapReader struct {
	metadata.MD
}

func (t TextMapReader) ForeachKey(handler func(key, val string) error) error {
	for key, val := range t.MD {
		for _, v := range val {
			if err := handler(key, v); err != nil {
				return err
			}
		}
	}
	return nil
}

//unary server interceptor
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !viper.GetBool("trace.isOpen") {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		tracer := opentracing.GlobalTracer()
		rootSpanCtx, err := tracer.Extract(opentracing.TextMap, TextMapReader{md})
		if err != nil && err != opentracing.ErrSpanContextNotFound {
			slog.Warn("tracer extract error: ", err.Error())
		}
		curSpan := tracer.StartSpan(
			info.FullMethod,
			ext.RPCServerOption(rootSpanCtx),
			ext.SpanKindRPCServer,
		)
		defer curSpan.Finish()

		//delivery span by context
		newCtx := opentracing.ContextWithSpan(ctx, curSpan)
		resp, err := handler(newCtx, req)
		if err == nil {
			if v, ok := resp.(*common.Response); ok {
				curSpan.SetTag("traceID", md["traceID"])
				curSpan.SetTag("resp.code", v.Code)
				curSpan.SetTag("resp.msg", v.Msg)
			}
		}

		return resp, err
	}
}
