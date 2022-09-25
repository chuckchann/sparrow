package mw_trace

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"sparrow/internal/pkg/slog"
	"sparrow/internal/pkg/util"
)

//unary client interceptor
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if !viper.GetBool("trace.isOpen") {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		var rootSpanCtx opentracing.SpanContext
		rootSpan := opentracing.SpanFromContext(ctx)
		if rootSpan != nil {
			rootSpanCtx = rootSpan.Context()
		}
		tracer := opentracing.GlobalTracer()
		span := tracer.StartSpan(
			method,
			opentracing.ChildOf(rootSpanCtx),
			ext.SpanKindRPCClient,
		)
		defer span.Finish()

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.New(nil)
		} else {
			md = md.Copy()
		}

		if md.Get("traceID") == nil {
			//get request traceID
			md.Set("traceID", util.GetTraceIDFromIncomingContext(ctx))
		}

		err := tracer.Inject(span.Context(), opentracing.TextMap, TextMapReader{md})
		if err != nil {
			slog.Error("tracer inject err, ", err.Error())
		}
		newCtx := metadata.NewOutgoingContext(ctx, md)

		return invoker(newCtx, method, req, reply, cc, opts...)
	}

}

/*
//stream grpc_client interceptor
func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(sctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {

	}



}

*/
