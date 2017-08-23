package models

var (
	DB   Database
	conn *gorm.DB
)

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

	conn.AutoMigrate(&Category{}, &Topic{})

	log.Info("AutMigrate database structs.")
}

func GetDB() *gorm.DB {
	if conn != nil {
		return conn
	}

	DB.Open()
	return conn
}
