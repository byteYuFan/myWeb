package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"myWeb/pkg/ttviper"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s\n", msg, err)
	}
}

var USERDB *gorm.DB
var (
	config        = ttviper.ConfigInit("config.yml")
	MysqlHost     = config.Viper.GetString("MySQL.Address")
	MysqlPort     = config.Viper.GetInt("MySQL.Port")
	MysqlUsername = config.Viper.GetString("MySQL.Username")
	MysqlPassword = config.Viper.GetString("MySQL.Password")
	MysqlDatabase = config.Viper.GetString("MySQL.Database")
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		MysqlUsername, MysqlPassword, MysqlHost, MysqlPort, MysqlDatabase,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	failOnError(err, "failed to connect database")
	sqlDB, err := db.DB()
	failOnError(err, "failed to get sql.DB")
	err = sqlDB.Ping()
	failOnError(err, "failed to ping mysql")
	USERDB = db
}
func Init() {

}
