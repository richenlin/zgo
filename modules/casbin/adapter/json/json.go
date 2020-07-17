package casbinjson

import (
	"io/ioutil"
	"github.com/suisrc/zgo/modules/casbin"
	"github.com/suisrc/zgo/modules/config"
	"github.com/suisrc/zgo/modules/files"

	"github.com/casbin/casbin/v2/persist"
	json "github.com/casbin/json-adapter/v2"

	"github.com/google/wire"
)

var _ persist.Adapter = (*json.Adapter)(nil)

// CasbinAdapterSet 注入casbin
var CasbinAdapterSet = wire.NewSet(casbin.NewCasbinEnforcer, NewCasbinAdapter)

// NewCasbinAdapter 构建Casbin Adapter
func NewCasbinAdapter() (persist.Adapter, error) {

	path := config.C.Casbin.PolicySource
	file, err := files.GetFile(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	// Initialize a JSON adapter and use it in a Casbin enforcer:
	adapter := json.NewAdapter(&data) // Use b as the data source.

	return adapter, nil
}

// rbac_policy.json
// [
//   {"PType":"p","V0":"alice","V1":"data1","V2":"read"},
//   {"PType":"p","V0":"bob","V1":"data2","V2":"write"},
//   {"PType":"p","V0":"data2_admin","V1":"data2","V2":"read"},
//   {"PType":"p","V0":"data2_admin","V1":"data2","V2":"write"},
//   {"PType":"g","V0":"alice","V1":"data2_admin"}
// ]
