package casbinredis

import (
	"zgo/modules/casbin"

	"github.com/casbin/casbin/v2/persist"
	redis "github.com/casbin/redis-adapter/v2"
	"github.com/google/wire"
)

var _ persist.Adapter = (*redis.Adapter)(nil)

// CasbinAdapterSet 注入casbin
var CasbinAdapterSet = wire.NewSet(casbin.NewCasbinEnforcer, NewCasbinAdapter)

// NewCasbinAdapter 构建Casbin Adapter
func NewCasbinAdapter() persist.Adapter {
	// Initialize a Redis adapter and use it in a Casbin enforcer:
	// key => casbin_rules
	adapter := redis.NewAdapter("tcp", "127.0.0.1:6379") // Your Redis network and address.
	return adapter
}
