package db

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	masterDB *gorm.DB
	slaveDB  *gorm.DB
)

func Init() {
	user := viper.GetString("db.rpc_service")
	password := viper.GetString("db.password")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	dbName := viper.GetString("db.name")
	maxOpenConns := viper.GetInt("db.maxOpenConns")
	maxIdleConns := viper.GetInt("db.maxOpenConns")

	args := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic("can not connect db, err:" + err.Error())
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetMaxIdleConns(maxIdleConns)

	err = db.DB().Ping()
	if err != nil {
		panic("can not ping db, err:" + err.Error())
	}

	masterDB = db
}

func GetMasterDB() *gorm.DB {
	return masterDB
}

func GetSlaveDB() *gorm.DB {
	return slaveDB
}
