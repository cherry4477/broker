package broker

import (
	"fmt"
)

type MysqlBroker struct {
}

func (mysql *MysqlBroker) Catalog() interface{} {
	fmt.Println("Mysql Catalog() called.")
	return mysql
}

func init() {
	register("mysql", &MysqlBroker{})
}
