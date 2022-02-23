package utility

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"time"
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

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	//fmt.Printf(format+"ee\n\rpp", args...)
	path := "/usr/src/myapp/gorm"
	writer, _ := rotatelogs.New(
		path+"%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(180)*time.Second),

		//这里设置1分钟产生一个日志文件
		rotatelogs.WithRotationTime(time.Duration(60)*time.Second),
	)
	//fmt.Println("aaa:" + format + "zzz")
	//fmt.Println("bbb:")
	//fmt.Println(args)
	//fmt.Println("ccc:")
	writer.Write([]byte(fmt.Sprintf(format, args...)))
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
	//baseLogPath := path.Join("log", "lg")
	//writer, err := rotatelogs.New(
	//	baseLogPath+".%Y%m%d%H%M",
	//	rotatelogs.WithLinkName(baseLogPath), // 生成软链，指向最新日志文件
	//
	//	rotatelogs.WithMaxAge(time.Hour*24*365), // 文件最大保存时间
	//	// rotatelogs.WithRotationCount(365),  // 最多存365个文件
	//
	//	rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
	//)

	newLogger := logger.New(
		Writer{},
		logger.Config{
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	//db.SingularTable(true)
	return

}
