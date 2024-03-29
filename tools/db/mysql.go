package db

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/wpf1118/toolbox/tools/flag"
	"github.com/wpf1118/toolbox/tools/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Mysql struct {
	*gorm.DB
}

var mysqldb *Mysql

func NewMysql() *Mysql {
	if mysqldb == nil {
		logging.ErrorF("db 模块没有初始化")
		return nil
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

	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	//	DefaultStringSize: 256, // string 类型字段的默认长度
	//	DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})
	//

	slowLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             5000 * time.Millisecond,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                           dsn,
		Conn:                          nil,
		SkipInitializeWithVersion:     false,
		DefaultStringSize:             100,
		DefaultDatetimePrecision:      nil,
		DisableDatetimePrecision:      false,
		DontSupportRenameIndex:        false,
		DontSupportRenameColumn:       false,
		DontSupportForShareClause:     false,
		DontSupportNullAsDefaultValue: false,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",                                // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                              // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false,                             // skip the snake_casing of names
			NameReplacer:  strings.NewReplacer("CID", "Cid"), // use name replacer to change struct/field name before convert it to db name
		},
		Logger: slowLogger,
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
