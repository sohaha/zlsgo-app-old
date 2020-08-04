package web

import (
	"fmt"
	"os"
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/ztime"
)

// Home Home 控制器
type Home struct{}

func (*Home) GetHome(c *znet.Context) {
	c.ApiJSON(200, "服务正常", map[string]interface{}{
		"time": ztime.Now(),
		"pid":  os.Getpid(),
	})
}

func (*Home) GetServerInfo(c *znet.Context) {
	if !c.Engine.IsDebug() {
		c.Log.Warn("非调试模式默认无法查看服务器信息")
		c.ApiJSON(403, "没有权限", nil)
	}
	const (
		B  = 1
		KB = 1024 * B
		MB = 1024 * KB
		GB = 1024 * MB
	)
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["mem"] = fmt.Sprintf("%.2fMB", float64(runtime.MemProfileRate)/MB)
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()

	dis, _ := disk.Usage("/")
	diskTotalGB := float64(dis.Total) / GB
	diskFreeGB := float64(dis.Free) / GB
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = fmt.Sprintf("%.2fGB", diskTotalGB)
	diskDic["free"] = fmt.Sprintf("%.2fGB", diskFreeGB)

	mem, _ := mem.VirtualMemory()
	memUsedMB := float64(mem.Used) / GB
	memTotalMB := float64(mem.Total) / GB
	memFreeMB := float64(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = fmt.Sprintf("%.2fGB", memTotalMB)
	memDic["used"] = fmt.Sprintf("%.2fGB", memUsedMB)
	memDic["free"] = fmt.Sprintf("%.2fGB", memFreeMB)
	memDic["usage"] = fmt.Sprintf("%d%%", memUsedPercent)
	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuNum"], _ = cpu.Counts(false)

	c.ApiJSON(200, "服务器信息", map[string]interface{}{
		"os":   osDic,
		"mem":  memDic,
		"cpu":  cpuDic,
		"disk": diskDic,
	})
}
