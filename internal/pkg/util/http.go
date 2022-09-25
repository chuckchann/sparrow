package util

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/parnurzeal/gorequest"
	"github.com/pkg/errors"
	"sparrow/internal/pkg/slog"
	"time"
)

type httpAgent struct {
	agent   *gorequest.SuperAgent
	params  interface{}
	header  map[string]string
	payload interface{}
}

func NewHttpAgent() *httpAgent {
	return &httpAgent{
		agent: gorequest.New(),
	}
}

type Option func(*httpAgent)

func WithParams(params interface{}) Option {
	return func(a *httpAgent) {
		a.params = params
	}
}

func WithHeaderMap(header map[string]string) Option {
	return func(a *httpAgent) {
		a.header = header
	}
}

func WithHeaderItems(items ...string) Option {
	return func(a *httpAgent) {
		if len(items)%2 != 0 {
			return
		}
		header := make(map[string]string, len(items)/2)
		for i := 0; i < len(items); i += 2 {
			header[items[i]] = items[i+1]
		}
		a.header = header
	}
}

func WithForm() Option {
	return func(a *httpAgent) {
		//todo

	}
}

func WithPayload(payload interface{}) Option {
	return func(a *httpAgent) {
		a.payload = payload
	}
}

func WithTrace(ctx context.Context) Option {
	return func(a *httpAgent) {
		curSpan := opentracing.SpanFromContext(ctx)
		err := opentracing.GlobalTracer().Inject(curSpan.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(a.agent.Header))
		if err != nil {
			slog.Warn("http agent inject failed ", err.Error())
		}
	}
}

type httpResponse struct {
	status int
	body   string
	err    error
}

func handleErrs(prefix string, errs []error) error {
	for _, e := range errs {
		if e != nil {
			return errors.Wrap(e, prefix)
		}
	}
	return nil
}

func recordHttpRequest(startTime time.Time, method string, url string, a *httpAgent) {
	slog.Info("method:", method, "url", url, "agent:", a, "startTime:", startTime, "cost:", time.Now().Sub(startTime))
}

func httpGet(url string, opts ...Option) httpResponse {
	a := NewHttpAgent()
	for _, opt := range opts {
		opt(a)
	}

	//defer recordHttpRequest(time.Now(), "GET", url, a)

	a.agent.Get(url).Query(a.params)
	for k, v := range a.header {
		a.agent.AppendHeader(k, v)
	}
	resp, body, errs := a.agent.End()
	if errs != nil {
		slog.Warn("HttpGet err: ", errs)
		return httpResponse{0, "", handleErrs("HttpGet:", errs)}
	}

	return httpResponse{resp.StatusCode, body, nil}
}

func HttpGet(ctx context.Context, url string, opts ...Option) (int, string, error) {
	ch := make(chan httpResponse, 1) //ch is a buffered channel, prevent goroutine leak
	go func() {
		ch <- httpGet(url, opts...)
	}()

	for {
		select {
		case resp := <-ch:
			return resp.status, resp.body, resp.err
		case <-ctx.Done():
			return 0, "", errors.New("http timeout")
		}
	}
}

func httpPost(url string, opts ...Option) httpResponse {
	a := NewHttpAgent()
	for _, opt := range opts {
		opt(a)
	}

	a.agent.Post(url).Send(a.payload)
	for k, v := range a.header {
		a.agent.AppendHeader(k, v)
	}
	resp, body, errs := a.agent.End()
	if errs != nil {
		slog.Warn("HttpGet err: ", errs)
		return httpResponse{0, "", handleErrs("HttpPost:", errs)}
	}

	return httpResponse{resp.StatusCode, body, nil}
}

func HttpPost(ctx context.Context, url string, opts ...Option) (int, string, error) {
	ch := make(chan httpResponse, 1) //ch is a buffered channel, prevent goroutine leak
	go func() {
		ch <- httpPost(url, opts...)
	}()

	for {
		select {
		case resp := <-ch:
			return resp.status, resp.body, resp.err
		case <-ctx.Done():
			return 0, "", errors.New("http timeout")
		}
	}
}
