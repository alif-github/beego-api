package main

import (
	"errors"
	"github.com/beego/i18n"
	"strings"
	_ "test_api/routers"
	"test_api/utils"
	"test_api/utils/databases"
	logs2 "test_api/utils/logs"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	var (
		errs    error
		runMode string
	)

	defer func() {
		if errs != nil {
			utils.ServerAttribute.Log.Error(errs.Error())
		}
	}()

	runMode, errs = beego.AppConfig.String("runmode")
	if errs != nil {
		return
	}

	if beego.BConfig.RunMode == runMode {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	} else {
		errs = errors.New("unset running mode")
		return
	}

	//--- Localize
	if errs = localize(); errs != nil {
		return
	}

	//--- Logs Set
	logs2.LogStore()

	//--- Database
	if errs = databases.SetDatabaseConnection(); errs != nil {
		return
	}

	//--- Logs Info Running
	utils.ServerAttribute.Log.Info("Running on localhost port 8080")

	//--- Run
	beego.Run()
}

func localize() (errs error) {
	var (
		langsKey = "langs"
		langs    string
	)

	//--- Set Localize
	_ = beego.AddFuncMap("i18n", i18n.Tr)

	//--- Localize
	langs, errs = beego.AppConfig.String(langsKey)
	if errs != nil {
		return
	}

	//--- Get Localize
	langsArr := strings.Split(langs, "|")
	for _, lang := range langsArr {
		if errs = i18n.SetMessage(lang, "conf/"+lang+".ini"); errs != nil {
			return
		}
	}

	return
}
