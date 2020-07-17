package middlewire

// 控制依赖的注入
import (
	"github.com/google/wire"
)

//===============================
// Default
//===============================

// DefaultGinSet 注入Injector
var DefaultGinSet = wire.NewSet(
	// wire.Struct(new(DefaultGinType), "*"),
	InitGinEngine, //*gin.Engine
	NewRouter,     // router
)

//===============================
// END
//===============================
