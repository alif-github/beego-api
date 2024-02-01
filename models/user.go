package models

import (
	"database/sql"
	"errors"
	"fmt"
	"test_api/dto/in"
	"test_api/utils"
)

var (
	UserList map[string]*in.User
)

func init() {
	UserList = make(map[string]*in.User)
	u := in.User{
		Id:       "user_11111",
		Username: "welcome",
		Password: "11111",
		Profile: in.Profile{
			Gender:  "male",
			Age:     20,
			Address: "Singapore",
			Email:   "astaxie@gmail.com",
		},
	}
	UserList["user_11111"] = &u
}

func AddUser(u *in.User) (errs error) {
	var (
		db  *sql.DB
		tx  *sql.Tx
		row *sql.Row
	)

	//--- Open Tx
	db = utils.ServerAttribute.Db
	tx, errs = db.Begin()
	if errs != nil {
		return
	}

	defer func() {
		if errs != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	query := fmt.Sprintf(`EXEC InsertOneUsers 
		@username = ?, 
		@password = ?, 
		@gender = ?, 
		@age = ?, 
		@address = ?, 
		@email = ? `)

	params := []interface{}{
		u.Username, u.Password, u.Profile.Gender,
		u.Profile.Age, u.Profile.Address, u.Profile.Email,
	}

	//--- Row
	row = tx.QueryRow(query, params...)
	if errs = row.Scan(&u.Id); errs != nil {
		return
	}

	return
}

func GetUser(uid int64) (u in.User, errs error) {
	query := fmt.Sprintf(`EXEC GetUserByID @id = ? `)
	params := []interface{}{uid}

	//--- Row
	row := utils.ServerAttribute.Db.QueryRow(query, params...)
	dbErr := row.Scan(&u.Id, &u.Username, &u.Password, &u.Profile.Age, &u.Profile.Gender, &u.Profile.Address, &u.Profile.Email, &u.Profile.Address)
	if dbErr != nil && dbErr.Error() != sql.ErrNoRows.Error() {
		return
	}

	return
}

func GetAllUsers() map[string]*in.User {
	return UserList
}

func UpdateUser(uid string, uu *in.User) (a *in.User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Address != "" {
			u.Profile.Address = uu.Profile.Address
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
