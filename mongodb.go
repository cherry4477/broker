package broker

import (
	"fmt"
)

type MongodbBroker struct{}

func (mongo *MongodbBroker) Catalog() interface{} {
	fmt.Println("MongoDB Catalog() called.")
	return mongo
}

func init() {
	register("mongodb", &MongodbBroker{})
}
