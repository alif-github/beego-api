package utils

import (
	"database/sql"
	"github.com/beego/beego/v2/core/logs"
)

var ServerAttribute serverAttribute

type serverAttribute struct {
	Log *logs.BeeLogger
	Db  *sql.DB
}
