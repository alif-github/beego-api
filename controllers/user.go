package controllers

import (
	"encoding/json"
	"github.com/beego/i18n"
	"net/http"
	"strconv"
	"test_api/dto"
	"test_api/dto/in"
	"test_api/models"
	"test_api/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
	i18n.Locale
}

// @Title CreateUser
// @Description create users
// @securityDefinitions.basic BasicAuth
// @Param body body models.User true "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var (
		user in.User
		errs error
		resp dto.Resp
	)

	defer func() {
		if errs != nil {
			utils.ServerAttribute.Log.Error(errs.Error())
			u.Ctx.Output.SetStatus(http.StatusInternalServerError)
			resp = dto.Resp{
				Code:    http.StatusInternalServerError,
				Message: errs.Error(),
			}
			u.Data["json"] = resp
			_ = u.ServeJSON()
		}
	}()

	if errs = json.Unmarshal(u.Ctx.Input.RequestBody, &user); errs != nil {
		return
	}

	errs = models.AddUser(&user)
	if errs != nil {
		return
	}

	//--- Resp JSON
	u.Data["json"] = dto.Resp{
		Code:    http.StatusOK,
		Message: "Success",
		Content: map[string]string{"id": user.Id},
	}

	_ = u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	var (
		errs error
		resp dto.Resp
	)

	defer func() {
		if r := recover(); r != nil || errs != nil {
			utils.ServerAttribute.Log.Error(errs.Error())
			u.Ctx.Output.SetStatus(http.StatusBadRequest)
			resp = dto.Resp{
				Code:    http.StatusBadRequest,
				Message: errs.Error(),
			}
			u.Data["json"] = resp
			_ = u.ServeJSON()
		}
	}()

	uid := u.GetString(":uid")
	id, errs := strconv.Atoi(uid)
	if errs != nil {
		return
	}

	var user in.User
	if id > 0 {
		user, errs = models.GetUser(int64(id))
		if errs != nil {
			return
		}
	}

	u.Data["json"] = dto.Resp{
		Code:    http.StatusOK,
		Message: "Success!!",
		Content: map[string]interface{}{
			"user": user,
		},
	}

	_ = u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user in.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if models.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
