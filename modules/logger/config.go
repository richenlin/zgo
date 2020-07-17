package logger

import (
	"context"
	"log/syslog"
	"os"
	"path/filepath"
	"zgo/modules/config"

	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
)

// InitLogger 初始化日志模块
func InitLogger(ctx context.Context) (func(), error) {
	c := config.C.Logging
	SetLevel(c.Level)
	SetFormatter(c.Format)

	// 设定日志输出
	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			SetOutput(os.Stdout)
		case "stderr":
			SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				_ = os.MkdirAll(filepath.Dir(name), 0777)

				f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				SetOutput(f)
				file = f
			}
		}
	}

	if c.EnableSyslogHook && c.SyslogAddr != "" {
		pri := syslog.LOG_INFO
		hook, err := logrus_syslog.NewSyslogHook(c.SyslogNetwork, c.SyslogAddr, pri, c.SyslogTag)
		if err != nil {
			Errorf(ctx, "Unable to connect to local syslog daemon")
		} else {
			AddHook(hook)
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}
