package orm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GormSetting struct {
	DBType    string
	Username  string
	Password  string
	Host      string
	DBName    string
	Charset   string
	ParseTime bool
	MaxIdleConns int
	MaxOpenConns int
}

func New(databaseSetting *GormSetting,debug bool) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.Username,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	)
	db,err:=gorm.Open(databaseSetting.DBType,connStr)
	if err != nil {
		return nil, err
	}
	if debug{
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)
	return db,nil
}
