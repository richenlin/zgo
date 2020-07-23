package endpoint

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/guonaihong/gout"

	goutapi "github.com/guonaihong/gout/interface"
	"github.com/suisrc/zgo/modules/auth/jwt"
)

// Config 配置
type Config struct {
	AddrURL    string                 // 服务器地址 http://abc.cn/api/auth/store
	Middleware goutapi.RequestMiddler // 可以增加远程调用规则和权限
}

// NewStore store
func NewStore(cfg *Config) *Store {
	client := &http.Client{}
	return &Store{
		cli:        client,
		url:        cfg.AddrURL,
		middleware: cfg.Middleware,
	}
}

var _ jwt.Storer = &Store{}

// Store redis存储
type Store struct {
	cli        *http.Client
	url        string
	middleware goutapi.RequestMiddler
}

// Set 存储令牌数据，并指定到期时间
func (s *Store) Set(ctx context.Context, tokenString string, expiration time.Duration) error {
	res := ResResult{}
	err := gout.New(s.cli).
		POST(s.url).
		SetJSON(gout.H{
			"token":   tokenString,
			"expired": expiration,
		}).
		RequestUse(s.middleware).
		BindJSON(&res).
		Do()

	if err != nil {
		return err
	}
	if !res.Success {
		return errors.New(res.ErrCode + "-" + res.ErrMessage)
	}
	return nil

}

// Check 检查令牌是否存在
func (s *Store) Check(ctx context.Context, tokenString string) (bool, error) {
	res := ResResult{}
	err := gout.New(s.cli).
		GET(s.url).
		SetQuery(gout.H{
			"token": tokenString,
		}).
		RequestUse(s.middleware).
		BindJSON(&res).
		Do()

	if err != nil {
		return false, err
	}
	if !res.Success {
		return false, errors.New(res.ErrCode + "-" + res.ErrMessage)
	}
	if result, ok := res.Data.(bool); ok {
		return result, nil
	}
	return false, errors.New("data is not bool")
}

// Close 关闭存储
func (s *Store) Close() error {
	s.cli.CloseIdleConnections()
	return nil
}

// ResResult 用于解析 服务端 返回的http body
type ResResult struct {
	Success    bool        `json:"success"`
	ErrMessage string      `json:"errmsg"`
	ErrCode    string      `json:"errcode"`
	Data       interface{} `json:"data"`
}
