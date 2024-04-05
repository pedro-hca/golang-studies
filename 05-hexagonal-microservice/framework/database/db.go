package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pedro-hca/golang-studies/05-rest-api/domain"
)

// type Database interface {
// 	Connect() (*gorm.DB, error)
// }

type DbConfig struct {
	DB     *gorm.DB
	Dsn    string
	DbType string
	Debug  bool
	Env    string
}

func NewDb() *DbConfig {
	return &DbConfig{}
}

// func NewDbTest() *gorm.DB {
// 	dbTest := NewDb()
// 	dbTest.DsnTest = ":memory:"
// 	dbTest.DbType = "sqlite3"
// 	dbTest.AutoMigrateDb = true
// 	dbTest.Debug = true
// 	dbTest.Env = "Test"
// }

func (db *DbConfig) Connect() (*gorm.DB, error) {

	dbConnection, err := gorm.Open(db.DbType, db.Dsn)
	if err != nil {
		return nil, err
	}

	dbConnection.LogMode(db.Debug)
	dbConnection.AutoMigrate(domain.User{})
	return dbConnection, nil

}
