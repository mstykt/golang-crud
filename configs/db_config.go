package configs

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang-crud/pkg/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

var DBConf dbConfig

type dbConfig struct {
	DB dbProps
}

type dbProps struct {
	Platform string
	Name     string
	Username string
	Password string
	Port     int
	Schema   string
}

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func (connection *Connection) GetSession() *gorm.DB {
	dbProps := DBConf.DB
	dns := "user=" + dbProps.Username + " password=" + dbProps.Password + " dbname=" + dbProps.Name + " port=" + strconv.Itoa(dbProps.Port) + " sslmode=disable TimeZone=Asia/Shanghai"
	sqlDB, err := sql.Open(dbProps.Platform, dns)
	if err != nil {
		panic(err)
	}

	db, gormErr := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbProps.Schema + ".",
			SingularTable: false,
		},
	})

	if gormErr != nil {
		panic(gormErr)
	}
	migrateErr := db.AutoMigrate(&entity.User{})
	if migrateErr != nil {
		panic(migrateErr)
	}
	return db
}
