package task

import (
	"github.com/sohaha/zlsgo/ztime"
	"github.com/sohaha/zlsgo/ztime/cron"

	"app/global"
)

// RegDemo 注册定时任务示例
func (*cornTask) RegDemo(c *cron.JobTable) {
	remove, _ := c.Add("* * * * *", func() {
		global.Log.Debug("每分钟执行一次:", ztime.Now())
	})
	// 移除定时任务
	remove()
}
