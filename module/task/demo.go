package task

import (
	"app/module"
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztime/cron"
)

// RegDemo 注册定时任务示例
func (*cornTask) RegDemo(c *cron.JobTable) {
	remove, _ := c.Add("* * * * *", func() {
		module.Log.Debug("每分钟执行一次:", ztime.Now())
	})
	// 移除定时任务
	remove()
}