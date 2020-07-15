package system

import "zgo/modules/logger"

// 服务器版本
var version = "0.0.1"

// SetVersion ver系统版本
func SetVersion(ver string) {
	version = ver
	logger.SetVersion(ver)
}

// 依赖主题版本(前端版本)
var requireThemeVersion = map[string][]string{
	"kratos-ui": {"0.0.1"},
}

// Version return the version of framework.
func Version() string {
	return version
}

// RequireThemeVersion return the require official version
func RequireThemeVersion() map[string][]string {
	return requireThemeVersion
}
