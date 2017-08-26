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
