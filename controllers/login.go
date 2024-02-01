package controllers

import (
	"encoding/json"
	"errors"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"test_api/dto"
	"test_api/dto/in"
	"test_api/models"
	"test_api/utils"
	"time"
)

// Operations about Login
type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description Login User
// @Param body body in.LoginModel true "body for login content"
// @Success 200 {object} in.LoginModel
// @router / [post]
func (l *LoginController) Login() {
	var (
		errs     error
		data     in.LoginModel
		input    in.LoginModel
		tokenStr string
		tokenMap = make(map[string]string)
	)

	defer func() {
		var resp dto.Resp
		if r := recover(); r != nil || errs != nil {
			utils.ServerAttribute.Log.Error(errs.Error())
			l.Ctx.Output.SetStatus(http.StatusUnauthorized)
			resp = dto.Resp{
				Code:    http.StatusUnauthorized,
				Message: errs.Error(),
			}
		} else {
			utils.ServerAttribute.Log.Info("Success Generated Token!!")
			l.Ctx.Output.SetStatus(http.StatusOK)
			resp = dto.Resp{
				Code:    http.StatusOK,
				Message: "Token has generated!",
				Content: tokenMap,
			}
		}

		//--- Give The Response
		l.Data["json"] = resp
		_ = l.ServeJSON()
	}()

	_ = json.Unmarshal(l.Ctx.Input.RequestBody, &input)
	data, errs = models.GetLoginInfo(input.Username)
	if errs != nil {
		errs = errors.New("Unauthorized!!!")
		return
	}

	if data.Password == input.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["msg"] = "This is success message from claims"
		claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

		tokenStr, errs = token.SignedString([]byte("oauth"))
		if errs != nil {
			return
		}

		tokenMap["token"] = tokenStr
	} else {
		errs = errors.New("Unauthorized!!!")
		return
	}
}
