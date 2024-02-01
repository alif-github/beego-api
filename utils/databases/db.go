package databases

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/denisenkom/go-mssqldb"
	"test_api/utils"
)

func SetDatabaseConnection() (errs error) {
	driver, _ := beego.AppConfig.String("database::driver")
	server, _ := beego.AppConfig.String("database::host")
	userID, _ := beego.AppConfig.String("database::user")
	password, _ := beego.AppConfig.String("database::password")
	port, _ := beego.AppConfig.String("database::port")
	database, _ := beego.AppConfig.String("database::database")
	schema, _ := beego.AppConfig.String("database::schema")

	utils.ServerAttribute.Db, errs = sql.Open("mssql", fmt.Sprintf(`%s:server=%s;user id=%s;password=%s;port=%s;database=%s;schema=%s;`,
		driver, server, userID, password, port, database, schema))
	if errs != nil {
		return
	}
	return errs
}
