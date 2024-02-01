package models

import (
	"database/sql"
	"fmt"
	"test_api/dto/in"
	"test_api/utils"
)

var (
	LoginList map[string]*in.LoginModel
)

func init() {
	LoginList = make(map[string]*in.LoginModel)
	l := in.LoginModel{
		Username: "admin",
		Password: "admin123",
	}
	LoginList["admin"] = &l
}

func GetLoginInfo(username string) (l in.LoginModel, errs error) {
	var db = utils.ServerAttribute.Db
	query := fmt.Sprintf(`EXEC GetPasswordUserByUsername @username = ? `)
	dbRow := db.QueryRow(query, username)
	if errs = dbRow.Scan(&l.Password); errs != nil && errs.Error() != sql.ErrNoRows.Error() {
		return
	}

	return
}
