package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"last-homework/model"
	"last-homework/tool"
)

// GDB mysql数据库的操作对象
var GDB *gorm.DB

func Mysql() (err error) {
	//获取配置
	user := tool.GetViper().GetString("mysql.user")
	pwd := tool.GetViper().GetString("mysql.password")
	h := tool.GetViper().GetString("mysql.host")
	p := tool.GetViper().GetString("mysql.port")
	db := tool.GetViper().GetString("mysql.dbname")
	dsn := user + ":" + pwd + "@tcp(" + h + ":" + p + ")/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
	gdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		return err
	}
	GDB = gdb

	err = createTables()
	return err
}

func createTables() error {
	if !GDB.Migrator().HasTable(&model.User{}) {
		err := GDB.AutoMigrate(&model.User{})
		if err != nil {
			return err
		}
	}
	if !GDB.Migrator().HasTable(&model.RoomMaker{}) {
		err := GDB.AutoMigrate(&model.RoomMaker{})
		if err != nil {
			return err
		}
	}

	return nil
}
