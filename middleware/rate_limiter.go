package middleware

import (
	"strconv"
	"time"
	"github.com/suisrc/zgo/modules/config"
	"github.com/suisrc/zgo/modules/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"golang.org/x/time/rate"
)

// RateLimiterMiddleware 请求频率限制中间件
func RateLimiterMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
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

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		if user, ok := helper.GetUserInfo(c); ok {
			limit := conf.Count
			rate, delay, allowed := limiter.AllowMinute(user.GetUserID(), limit)
			if !allowed {
				h := c.Writer.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(limit-rate, 10))
				delaySec := int64(delay / time.Second)
				h.Set("X-RateLimit-Delay", strconv.FormatInt(delaySec, 10))
				helper.ResError(c, &helper.Err429TooManyRequests)
				return
			}
		}

		c.Next()
	}
}
