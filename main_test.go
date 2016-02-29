package broker

import "testing"

func TestBroker(t *testing.T) {
	mysql, _ := New("mysql")
	mongo, _ := New("mongodb")

	mysql.Catalog()
	mysql.Binding()

	mongo.Catalog()
	mongo.Provisioning()
}

/*

package main

import (
	"github.com/zonesan/broker"
)

func main() {
	mysql, _ := broker.New("mysql")
	mongo, _ := broker.New("mongodb")

	mysql.Catalog()
	mysql.Binding()

	mongo.Catalog()
	mongo.Provisioning()
}

}

*/
