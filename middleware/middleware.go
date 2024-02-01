package middleware

import (
	"errors"
	"github.com/beego/beego/v2/adapter/plugins/cors"
	context2 "github.com/beego/beego/v2/server/web/context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"test_api/dto"
)

func Middleware(ctx *context2.Context) {
	cors.Allow(&cors.Options{
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "content-type", "Access-Control-Allow-Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	})
}

func AuthMiddleware(ctx *context2.Context) {
	var (
		modelRespAuth dto.Resp
		tokenJWT      *jwt.Token
		errs          error
		secretKey     = "oauth"
		defaultErrMsg = "Unauthorized: Missing Token!"
	)

	defer func() {
		if errs != nil {
			modelRespAuth = dto.Resp{
				Code:    http.StatusUnauthorized,
				Message: errs.Error(),
			}

			//--- Send To Response
			ctx.Output.SetStatus(http.StatusUnauthorized)
			_ = ctx.JSONResp(modelRespAuth)
		}
	}()

	token := ctx.Input.Header("Authorization")
	if token == "" {
		errs = errors.New(defaultErrMsg)
		return
	}

	tokenJWT, errs = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if errs != nil || !tokenJWT.Valid {
		errs = errors.New(defaultErrMsg)
		return
	}

	//--- Set the user information
	claims, _ := tokenJWT.Claims.(jwt.MapClaims)
	ctx.Input.SetData("claims", claims)
}
