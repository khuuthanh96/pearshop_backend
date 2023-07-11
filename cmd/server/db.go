package main

import (
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"

	"pearshop_backend/app/config"
	"pearshop_backend/pkg/gormutil"
	"pearshop_backend/pkg/log"
)

func initDBConnection(cfg config.MySQL) {
	if cfg.Masters == "" {
		log.Errorln("miss db info")
		os.Exit(1)
	}

	var (
		db         *gorm.DB
		gormConfig = gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
		cfgMasters       = strings.Split(cfg.Masters, ",")
		cfgSlaves        = strings.Split(cfg.Slaves, ",")
		masterDialectors = make([]gorm.Dialector, 0, len(cfgMasters))
		slaveDialectors  = make([]gorm.Dialector, 0, len(cfgSlaves))
	)

	if cfg.IsEnabledLog {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	master, otherMasters := cfgMasters[0], cfgMasters[1:]

	for _, otherMaster := range otherMasters {
		if otherMaster == "" {
			continue
		}

		dialector := mysql.New(mysql.Config{DSN: cfg.Conn(otherMaster)})
		masterDialectors = append(masterDialectors, dialector)
	}

	for _, slave := range cfgSlaves {
		if slave == "" {
			continue
		}

		dialector := mysql.New(mysql.Config{DSN: cfg.Conn(slave)})
		slaveDialectors = append(slaveDialectors, dialector)
	}

	db, err := gormutil.OpenDBConnection(cfg.Conn(master), gormConfig)
	if err != nil {
		log.WithError(err).
			Errorln("creating connection to DB")
		os.Exit(1)
	}

	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  masterDialectors,
		Replicas: slaveDialectors,
		Policy:   dbresolver.RandomPolicy{},
	})); err != nil {
		log.WithError(err).
			WithField("master_dialectors", masterDialectors).
			WithField("slave_dialectors", slaveDialectors).
			Errorln("fail to register master slave dbs")
		os.Exit(1)
	}

	rawDB, err := db.DB()
	if err != nil {
		log.WithError(err).Errorln("get DB failed")
		os.Exit(1)
	}

	rawDB.SetMaxOpenConns(cfg.MaxOpenConns)
	rawDB.SetMaxIdleConns(cfg.MaxIdleConns)
	rawDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Minute)
}
