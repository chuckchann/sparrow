package mw_rate_limit

import (
	"context"
	"github.com/spf13/viper"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"sparrow/internal/pkg/serror"
	"sync"
	"time"
)

var limiter RateLimiter

func init() {
	limiter = rate.NewLimiter(50, 100)
}

type RateLimiter interface {
	Allow() bool
}

type BucketLimiter struct {
	rate      int64 //
	size      int64 //
	timestamp int64 //unix timestamp
	cur       int64 //current water
	mu        sync.Mutex
}

func NewBucketLimiter(rate int64, size int64) *BucketLimiter {
	return &BucketLimiter{
		rate:      rate,
		size:      size,
		timestamp: time.Now().Unix(),
		cur:       0,
	}
}

func (b *BucketLimiter) Allow() bool {
	nowTimestamp := time.Now().Unix()
	var allowed bool

	b.mu.Lock()
	defer b.mu.Unlock()
	count := (nowTimestamp - b.timestamp) * b.rate
	if count < 0 {
		b.cur = 0
		allowed = true
	} else {
		b.cur = count
		allowed = b.cur > b.size
	}
	b.timestamp = nowTimestamp
	return allowed
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !viper.GetBool("rateLimit.isOpen") {
			return handler(ctx, req)
		}
		if !limiter.Allow() {
			return serror.ERR_RATE.Response(), nil
		}
		return handler(ctx, req)
	}
}
