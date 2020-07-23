// +build document

package xorm

/*
 不推荐使用xorm或者gorm
*/
import "xorm.io/xorm"

// Demo demo
func Demo() {
	xorm.NewEngine("sqlite3", "./test.db")
}
