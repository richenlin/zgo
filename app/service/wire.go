package service

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/suisrc/zgo/app/model/ent"
	"github.com/suisrc/zgo/app/model/entc"
	"github.com/suisrc/zgo/app/model/sqlxc"
)

// ServiceSet wire注入服务
var ServiceSet = wire.NewSet(
	// 数据库连接注册
	entc.NewClient,
	sqlxc.NewClient,
	wire.Struct(new(GPA), "*"),
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

// GPA golang persistence api 数据持久化
type GPA struct {
	DBE *ent.Client // ent client
	DBS *sqlx.DB    // sqlx client
}
