package casbinsqlx

import (
	"errors"

	"github.com/suisrc/zgo/modules/casbin"
	"github.com/suisrc/zgo/modules/config"

	"github.com/casbin/casbin/v2/persist"
	"github.com/google/wire"
	sqlx "github.com/memwey/casbin-sqlx-adapter"
)

var _ persist.Adapter = (*sqlx.Adapter)(nil)

// CasbinAdapterSet 注入casbin
var CasbinAdapterSet = wire.NewSet(casbin.NewCasbinEnforcer, NewCasbinAdapter)

// NewCasbinAdapter 构建Casbin Adapter
func NewCasbinAdapter() (persist.Adapter, error) {

	pType := config.C.Casbin.PolicyType
	pSource := config.C.Casbin.PolicySource
	pTable := config.C.Casbin.PolicyTable
	if pType == "" || pSource == "" {
		return nil, errors.New("Casbin.PlicyType OR Casbin.PolicySource has Empty")
	}
	if pTable == "" {
		pTable = "casbin_rule"
	}

	opts := &sqlx.AdapterOptions{
		DriverName:     pType,
		DataSourceName: pSource,
		TableName:      pTable,
		// or reuse an existing connection:
		// DB: myDBConn,
	}
	adapter := sqlx.NewAdapterFromOptions(opts)
	return adapter, nil
}
