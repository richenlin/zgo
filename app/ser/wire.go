package ser

import (
	"context"
	"log"

	"github.com/google/wire"
	"github.com/suisrc/zgo/app/ent"

	// 引入数据库
	_ "github.com/mattn/go-sqlite3"
)

// ServiceSet wire注入声明
var ServiceSet = wire.NewSet(
	NewDBClient,

	// 服务注册
	wire.Struct(new(User), "*"),
)

// NewDBClient client
func NewDBClient() (*ent.Client, func(), error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}

	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// defer client.Close()
	clean := func() {
		client.Close()
	}
	return client, clean, nil
}
