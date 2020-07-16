package injector

import (
	"zgo/router"

	"github.com/google/wire"
)

// InitRoutesSet 注入到wire中
var InitRoutesSet = wire.NewSet(wire.Struct(new(InitRoutesOptions), "*"), InitRoutesFunc)

// InitRoutesOptions options
type InitRoutesOptions struct {
}

// InitRoutesResult result
type InitRoutesResult struct {
	root *router.RootPath
}

// InitRoutesFunc func
func InitRoutesFunc(opts *InitRoutesOptions) *InitRoutesResult {
	return &InitRoutesResult{
		root: nil,
	}
}
