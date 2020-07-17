package casbin

import (
	"errors"
	"time"

	"github.com/suisrc/zgo/modules/logger"

	"github.com/suisrc/zgo/modules/config"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
)

// NewCasbinEnforcer 初始化casbin enforcer
func NewCasbinEnforcer(adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	c := config.C.Casbin
	if c.Model == "" {
		// return new(casbin.SyncedEnforcer), func() {}, nil
		return nil, nil, errors.New("Casbin Model no config")
	}

	enforcer, err := casbin.NewSyncedEnforcer(c.Model)
	if err != nil {
		return nil, nil, err
	}
	logger.Infof(nil, "Loading Casbin Model[%s]", c.Model)
	enforcer.EnableLog(c.Debug)

	err = enforcer.InitWithModelAndAdapter(enforcer.GetModel(), adapter)
	if err != nil {
		return nil, nil, err
	}
	enforcer.EnableEnforce(c.Enable)

	cleanFunc := func() {}
	if c.AutoLoad {
		enforcer.StartAutoLoadPolicy(time.Duration(c.AutoLoadInternal) * time.Second)
		cleanFunc = func() {
			enforcer.StopAutoLoadPolicy()
		}
	}

	return enforcer, cleanFunc, nil
}
