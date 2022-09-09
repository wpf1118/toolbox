package db

import (
	"fmt"
	"github.com/wpf1118/toolbox/tools/flag"
	"github.com/wpf1118/toolbox/tools/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
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
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",                                // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                             // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
	})

	if err != nil {
		err = logging.ErrorF("could not connect to the database: %v", err)
		return nil, err
	}

	return &Mysql{
		db,
	}, nil
}

func (m *Mysql) List() {

}
