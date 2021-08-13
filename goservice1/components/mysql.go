package components

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func mysqlCreateClient() *gorm.DB {
	//初始化Db
	db, dberr := gorm.Open(
		mysql.Open(Config.Section("mysql").Key("dsn").String()),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "ly_", // 表名前缀，`User` 的表名应该是 `t_users`
				SingularTable: true,  // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
		},
	)
	if dberr != nil {
		panic("mysql error")
	}

	//defer db.Close()
	return db
}
