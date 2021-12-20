// Created at 11/29/2021 10:04 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package orm_manager

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm"
)

type Manager struct {
	clients sync.Map
}

var manager = &Manager{}

func InstallManager(configs []MySqlConfig) {
	for _, config := range configs {
		db, err := config.initClient()
		if err != nil {
			panic(err)
		}

		if val, ok := manager.clients.Load(config.Name); ok {
			err := fmt.Errorf("InstallManager - err: dublicated config name {%v}", val)
			panic(err)
		}
		manager.clients.Store(config.Name, db)
	}
}

func GetClient(dbName string) (db *gorm.DB) {
	if val, ok := manager.clients.Load(dbName); ok {
		if db, ok = val.(*gorm.DB); ok {
			return
		}
	}

	log.Printf("orm_manager::GetClient - Not found client: %s", dbName)
	return
}
