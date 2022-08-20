package db

import (
	"fmt"
	"github.com/wpf1118/toolbox/tools/flag"
	"github.com/wpf1118/toolbox/tools/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	*gorm.DB
}

var mysqldb *Mysql

func NewMysql() *Mysql {
	if mysqldb == nil {
		logging.ErrorF("db 模块没有初始化")
		panic("db 模块没有初始化")
	}

	return mysqldb
}

func MysqlInit(mysqlOpts *flag.MysqlOpts) {
	var err error
	mysqldb, err = newMysqlClient(mysqlOpts)
	if err != nil {
		logging.ErrorF("NewMysqlClient error: %v", err)
	}
}

func newMysqlClient(mysqlOpts *flag.MysqlOpts) (*Mysql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&autocommit=true",
		mysqlOpts.Username,
		mysqlOpts.Password,
		mysqlOpts.Endpoint,
		mysqlOpts.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		err = logging.ErrorF("could not connect to the database: %v", err)
		return nil, err
	}

	// eg
	type Kv struct {
		ID        uint           `json:"id" gorm:"primarykey"`
		Key       string         `json:"key"`
		Value     string         `json:"value"`
		CreatedAt int64          `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt int64          `json:"updated_at" gorm:"autoUpdateTime"`
		DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	}

	db.AutoMigrate(&Kv{})

	return &Mysql{
		db,
	}, nil
}
