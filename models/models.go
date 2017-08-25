/*
Copyright 2014 Huawei Technologies Co., Ltd. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"os"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"

	"github.com/xiechuanj/blog/configure"
)

var (
	// DB is
	DB   Database
	conn *gorm.DB
)

// Database is
type Database interface {
	Open()
	Migrate()
}

type defaultDB struct {
}

func init() {
	if DB == nil {
		DB = new(defaultDB)
	}
}

func (d *defaultDB) Open() {
	var err error
	if conn, err = gorm.Open(configure.GetString("database.driver"), configure.GetString("database.uri")); err != nil {
		log.Fatal("Initlization database connection error.")
		os.Exit(1)
	} else {
		conn.DB()
		conn.DB().Ping()
		conn.DB().SetMaxIdleConns(10)
		conn.DB().SetMaxOpenConns(100)
		conn.SingularTable(true)
	}
}

func (d *defaultDB) Migrate() {
	d.Open()
	conn.AutoMigrate(&Topic{}, &Category{})

	log.Info("AutMigrate database structs.")
}

// GetDB is
func GetDB() *gorm.DB {
	if conn != nil {
		return conn
	}

	DB.Open()
	return conn
}
