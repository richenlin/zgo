package casbincustom

import (
	"zgo/modules/casbin"

	"github.com/casbin/casbin/v2/persist"

	"github.com/google/wire"
)

var _ persist.Adapter = (*Adapter)(nil)

// CasbinAdapterSet 注入casbin
var CasbinAdapterSet = wire.NewSet(casbin.NewCasbinEnforcer, NewCasbinAdapter)

// NewCasbinAdapter 构建Casbin Adapter
func NewCasbinAdapter() persist.Adapter {
	// Initialize a JSON adapter and use it in a Casbin enforcer:

	b := []byte{}             // b stores Casbin policy in JSON bytes.
	adapter := NewAdapter(&b) // Use b as the data source.

	return adapter
}

// rbac_policy.json
// [
//   {"PType":"p","V0":"alice","V1":"data1","V2":"read"},
//   {"PType":"p","V0":"bob","V1":"data2","V2":"write"},
//   {"PType":"p","V0":"data2_admin","V1":"data2","V2":"read"},
//   {"PType":"p","V0":"data2_admin","V1":"data2","V2":"write"},
//   {"PType":"g","V0":"alice","V1":"data2_admin"}
// ]
