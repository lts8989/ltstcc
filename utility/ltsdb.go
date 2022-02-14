package utility

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var usersDb *gorm.DB
var goodsDb *gorm.DB
var ordersDb *gorm.DB

type dbConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsersDb() *gorm.DB {
	if usersDb == nil {
		usersDb = getConfig("userDb.yaml")
	}
	return usersDb
}

func GetGoodsDb() *gorm.DB {
	if goodsDb == nil {
		goodsDb = getConfig("goodDb.yaml")
	}
	return goodsDb
}

func GetOrdersDb() *gorm.DB {
	if ordersDb == nil {
		ordersDb = getConfig("orderDb.yaml")
	}
	return ordersDb
}

func getConfig(configFileName string) (db *gorm.DB) {
	var setting dbConfig
	config, err := ioutil.ReadFile("./" + configFileName)
	if err != nil {
		fmt.Println(err)
	}
	err = yaml.Unmarshal(config, &setting)
	if err != nil {
		fmt.Println(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", setting.Username, setting.Password, setting.Host, setting.Port, setting.Dbname)
	fmt.Println(configFileName + dsn)
	db, err = gorm.Open("mysql", dsn)
	db.SingularTable(true)
	return

}
