// +build document

package sqlxc

/*
 注意,该代码不会进行编译
*/
import (
	"database/sql"
	"time"

	// 引入数据库
	_ "github.com/mattn/go-sqlite3"
)

// NewClient client
func NewClient() (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite3", "file:db1?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}
	db.SetMaxIdleConns(10)           // 最大空闲连接数
	db.SetMaxOpenConns(100)          // 数据库最大连接数
	db.SetConnMaxLifetime(time.Hour) //连接最长存活期，超过这个时间连接将不再被复用

	// defer client.Close()
	clean := func() {
		db.Close()
	}
	return db, clean, nil
}
