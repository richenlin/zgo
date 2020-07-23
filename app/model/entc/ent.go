package entc

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/suisrc/zgo/app/model"
	"github.com/suisrc/zgo/app/model/ent"

	// 引入数据库
	_ "github.com/mattn/go-sqlite3"
)

// NewClient client
func NewClient() (*ent.Client, func(), error) {

	drv, err := sql.Open("sqlite3", "file:db1?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	client := ent.NewClient(ent.Driver(drv))
	//client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	//if err != nil {
	//	return nil, nil, err
	//}

	// run the auto migration tool.
	if model.TableSchemaInitEnt || model.TableSchemaInit {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		} else {
			// 防止其他持久化框架更新table结构
			model.TableSchemaInit = false
		}
	}

	// defer client.Close()
	clean := func() {
		client.Close()
	}
	return client, clean, nil
}
