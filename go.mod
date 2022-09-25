module sparrow

go 1.14

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/elazarl/goproxy v0.0.0-20210110162100-a92cc753f88e // indirect
	github.com/fsnotify/fsnotify v1.4.7
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/jinzhu/gorm v1.9.16
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/parnurzeal/gorequest v0.2.16
	github.com/pborman/uuid v1.2.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.9.0
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.7.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	golang.org/x/tools v0.1.0 // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/redis.v5 v5.2.9
	moul.io/http2curl v1.0.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
