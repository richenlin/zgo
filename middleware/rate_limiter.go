package middleware

import (
	"strconv"
	"time"
	"zgo/engine"
	"zgo/modules/config"
	"zgo/modules/result"

	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"golang.org/x/time/rate"
)

// RateLimiterMiddleware 请求频率限制中间件
func RateLimiterMiddleware(skippers ...SkipperFunc) engine.HandlerFunc {
	conf := config.C.RateLimiter
	if !conf.Enable {
		return EmptyMiddleware()
	}

	rc := config.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       conf.RedisDB,
	})

	limiter := redis_rate.NewLimiter(ring)
	limiter.Fallback = rate.NewLimiter(rate.Inf, 0)

	return func(c engine.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		if user, ok := c.GetUserInfo(); ok {
			limit := conf.Count
			rate, delay, allowed := limiter.AllowMinute(user.GetUserID(), limit)
			if !allowed {
				h := c.ResponseWriter().Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(limit-rate, 10))
				delaySec := int64(delay / time.Second)
				h.Set("X-RateLimit-Delay", strconv.FormatInt(delaySec, 10))
				result.ResError(c, result.Err429TooManyRequests)
				return
			}
		}

		c.Next()
	}
}
