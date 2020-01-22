package worker

import (
	"github.com/robfig/cron"
	common "github.com/tsbxmw/gin_common"
)

func CornWork() {
	common.LogrusLogger.Info("Corn work start")
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/5 * * * * *", MessageSendCheckWork)
	c.Start()
	select {}
}

func MessageSendCheckWork() {
	common.LogrusLogger.Info("message send check work start")
}
