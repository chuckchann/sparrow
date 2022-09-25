package trace

import (
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
)

func Init() {
	cfg := jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const", //固定采样
			Param: 1,       //1=全采样、0=不采样
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
		ServiceName: viper.GetString("appName"),
	}

	//closer 在后面回收时有用 这里先不处理
	tracer, _, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	if err != nil {
		errMsg := "cannot init jaeger, err: " + err.Error()
		panic(errMsg)
	}

	opentracing.SetGlobalTracer(tracer)
}
