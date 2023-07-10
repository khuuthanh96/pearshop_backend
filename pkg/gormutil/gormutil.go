package gormutil

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	once      sync.Once
	singleton *gorm.DB
)

// OpenDBConnection opens a DB connection.
func OpenDBConnection(conn string, config gorm.Config) (*gorm.DB, error) {
	var err error

	once.Do(func() {
		singleton, err = gorm.Open(mysql.Open(conn), &config)
		if err != nil {
			err = fmt.Errorf("gorm open: %w", err)

			return
		}
	})

	if err != nil {
		return nil, err
	}

	return singleton, nil
}

// GetDB gets the instance of gorm singleton
func GetDB() *gorm.DB {
	return singleton
}
