package logs

import (
	"github.com/beego/beego/v2/adapter/logs"
	log2 "github.com/beego/beego/v2/core/logs"
	"test_api/utils"
)

func LogStore() {
	var channelLens = int64(1000)
	log := log2.NewLogger(channelLens)
	_ = log.SetLogger(logs.AdapterConsole, `{"color":true}`)
	utils.ServerAttribute.Log = log
}
