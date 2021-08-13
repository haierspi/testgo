package components

import (
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var (
	Config *ini.File
	Db     *gorm.DB
	Redis  *redis.Client
)

func Init() {

	var iniErr error
	Config, iniErr = ini.Load("config/config.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	//创建grom db 连接
	Db = mysqlCreateClient()
	//创建
	Redis = redisCreateClient()

}
