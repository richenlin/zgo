package ser

import (
	"github.com/google/wire"
	"github.com/suisrc/zgo/app/model/entc"
	"github.com/suisrc/zgo/app/model/sqlxc"
)

// ServiceSet wire注入服务
var ServiceSet = wire.NewSet(
	// 数据库连接注册
	entc.NewClient,
	sqlxc.NewClient,
	// 服务
	wire.Struct(new(Demo), "*"),
)

//======================================
// 分割线
//======================================

// ResultRef 返回值暂存器
type ResultRef struct {
	Data interface{}
}
