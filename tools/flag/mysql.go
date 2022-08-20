package flag

import (
	"github.com/wpf1118/toolbox/tools/env"
)

// MysqlOpts the Mongo options.
type MysqlOpts struct {
	Endpoint string
	Username string
	Password string
	Database string
}

// NewDefaultMysqlOpts returns a new default mongodb options.
func NewDefaultMysqlOpts() *MysqlOpts {
	return &MysqlOpts{
		Endpoint: env.GetEnv(env.MysqlEndpoint, "localhost:3306"),
		Username: env.GetEnv(env.MysqlUsername, "root"),
		Password: env.GetEnv(env.MysqlPassword, "root"),
		Database: env.GetEnv(env.MysqlDatabase, "api"),
	}
}
