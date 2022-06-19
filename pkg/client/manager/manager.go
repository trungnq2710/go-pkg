package manager

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"
)

var ins sync.Map

func Install(configs []Config) {
	for _, config := range configs {
		db, err := config.InitClient()
		if err != nil {
			panic(err)
		}
		if val, ok := ins.Load(config.GetConfigName()); ok {
			err := fmt.Errorf("manager::Install - err: duplicated config name {%v}", val)
			panic(err)
		}
		ins.Store(config.GetConfigName(), db)
	}
}

func GetClient(name string) interface{} {
	if db, ok := ins.Load(name); ok {
		return db
	}
	log.Printf("manager::GetClient - Not found client: %s", name)
	return nil
}

func GetMongoClient(name string) (db *mongo.Database) {
	if val, ok := ins.Load(name); ok {
		if db, ok = val.(*mongo.Database); ok {
			return
		}
	}
	log.Printf("manager::GetMongoClient - Not found client: %s", name)
	return
}
