// Created at 11/29/2021 10:04 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package db_manager

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"sync"

	"gorm.io/gorm"
)

type Manager struct {
	clients sync.Map
}

var manager = &Manager{}

func InstallManager(configs []Config) {
	for _, config := range configs {
		db, err := config.InitClient()
		if err != nil {
			panic(err)
		}

		if val, ok := manager.clients.Load(config.GetName()); ok {
			err := fmt.Errorf("InstallManager - err: dublicated config name {%v}", val)
			panic(err)
		}
		manager.clients.Store(config.GetName(), db)
	}
}

func GetClient(dbName string) interface{} {
	if db, ok := manager.clients.Load(dbName); ok {
		return db
	}

	log.Printf("orm_manager::GetClient - Not found client: %s", dbName)
	return nil
}

func GetGormClient(dbName string) (db *gorm.DB) {
	if val, ok := manager.clients.Load(dbName); ok {
		if db, ok = val.(*gorm.DB); ok {
			return
		}
	}

	log.Printf("orm_manager::GetGormClient - Not found client: %s", dbName)
	return
}

func GetMongoClient(dbName string) (db *mongo.Database) {
	if val, ok := manager.clients.Load(dbName); ok {
		if db, ok = val.(*mongo.Database); ok {
			return
		}
	}

	log.Printf("orm_manager::GetMongoClient - Not found client: %s", dbName)
	return
}
