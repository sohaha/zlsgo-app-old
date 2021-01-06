package logic

import (
	"runtime"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"

	"github.com/sohaha/zlsgo/zfile"
)

func GetServerInfo() map[string]interface{} {
	osDic := make(map[string]interface{}, 0)
	osDic["goOs"] = runtime.GOOS
	osDic["arch"] = runtime.GOARCH
	osDic["memory"] = zfile.SizeFormat(uint64(runtime.MemProfileRate))
	osDic["compiler"] = runtime.Compiler
	osDic["version"] = runtime.Version()
	osDic["numGoroutine"] = runtime.NumGoroutine()

	dis, _ := disk.Usage("/")
	diskDic := make(map[string]interface{}, 0)
	diskDic["total"] = zfile.SizeFormat(dis.Total)
	diskDic["free"] = zfile.SizeFormat(dis.Free)

	memory, _ := mem.VirtualMemory()
	memDic := make(map[string]interface{}, 0)
	memDic["total"] = zfile.SizeFormat(memory.Total)
	memDic["used"] = zfile.SizeFormat(memory.Used)
	memDic["free"] = zfile.SizeFormat(memory.Free)
	memDic["usage"] = int(memory.UsedPercent)
	cpuDic := make(map[string]interface{}, 0)
	cpuDic["cpuNum"], _ = cpu.Counts(false)
	return map[string]interface{}{
		"os":     osDic,
		"memory": memDic,
		"cpu":    cpuDic,
		"disk":   diskDic,
	}
}
