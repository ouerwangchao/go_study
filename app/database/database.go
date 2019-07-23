package database

import (
	//import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"sync"
	"time"
)

var (
	MasterDB *gorm.DB
)

//Init func init MySQL db connection pool
//TODO: use Init(viper.Sub("mysql")) func
func Init() {
	var once sync.Once
	initDB := func() {
		db, err := gorm.Open("mysql",
			viper.GetString("mysql.master.user")+":"+
				viper.GetString("mysql.master.password")+"@tcp("+
				viper.GetString("mysql.master.host")+":"+
				viper.GetString("mysql.master.port")+")/"+
				viper.GetString("mysql.master.db")+"?charset"+
				viper.GetString("mysql.master.charset")+"&parseTime=True&loc=Local")
		if err != nil {
			panic(err)
		}
		MasterDB = db
		MasterDB.DB().SetMaxOpenConns(viper.GetInt("mysql.master.max_open_conns"))
		MasterDB.DB().SetMaxIdleConns(viper.GetInt("mysql.master.max_idle_conns"))
		MasterDB.DB().SetConnMaxLifetime(time.Minute * time.Duration(viper.GetInt("mysql.mas)ter.conn_max_lifetime")))
		MasterDB.LogMode(viper.GetBool("mysql.master.log_mode"))
	}
	if MasterDB == nil {
		//TODO: add log
		once.Do(initDB)
	} else if err := MasterDB.DB().Ping(); err != nil {
		//TODO: add log
		initDB()
	} else if err := MasterDB.Error; err != nil {
		//TODO: add log
		panic(err)
	}
}
