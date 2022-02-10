// Created at 11/29/2021 10:12 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package db_manager

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlConfig struct {
	Name     string
	DSN      string
	LogLevel int
}

func (m *MySqlConfig) GetName() string {
	return m.Name
}

func (m *MySqlConfig) Valid() (err error) {
	if m.Name == "" {
		err = fmt.Errorf("config name is empty")
		return
	}
	if m.DSN == "" {
		err = fmt.Errorf("config DNS is empty")
		return
	}
	return
}

func (m *MySqlConfig) InitClient() (interface{}, error) {
	if err := m.Valid(); err != nil {
		return nil, err
	}
	return gorm.Open(mysql.Open(m.DSN), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.LogLevel(m.LogLevel)),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}
