package manage

import (
	"app/global"
	"app/logic"
	"app/model"
	"app/web/business/manageBusiness"
	"fmt"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"
	"os"
	"strings"
	"time"
)

// 后台-系统接口
type System struct {
}

// GetSystemLogs 系统日志
func (*System) GetSystemLogs(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var postData manageBusiness.GetSystemLogsSt
	tempRule := c.ValidRule()
	err := zvalid.Batch(
		zvalid.BatchVar(&postData.Name, c.Valid(tempRule, "name", "文件名称")),
		zvalid.BatchVar(&postData.Type, c.Valid(tempRule, "type", "文件夹名称")),
		zvalid.BatchVar(&postData.CurrentLine, c.Valid(tempRule, "currentLine", "当前行")),
	)
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	showDir := manageBusiness.GetTmpLogDir(global.BaseConf().LogDir)
	logLists := manageBusiness.ShowLogsLists(postData.Type, global.BaseConf().LogDir)

	logPath := zfile.RealPath("./" + global.BaseConf().LogDir + "/" + postData.Type + "/" + postData.Name)
	var (
		fileContentSlice []string
		fileSize         int64
		fileLine         int
	)
	if global.BaseConf().LogDir != "" {
		fileContentSlice, fileSize, fileLine = manageBusiness.GetSystemLogInfo(logPath, postData.CurrentLine)
	}

	fileContent := ""
	if fileSize > (1024 * 500) {
		fileContent = "日志文件过大（" + fmt.Sprintf("%v", logic.Byte2Kb(fileSize)) + "kb），不支持在线查看全部内容。\n\n"
		fileContent += strings.Join(fileContentSlice[len(fileContentSlice)-10:], "")
	} else {
		fileContent = strings.Join(fileContentSlice, "")
	}

	c.ApiJSON(200, "系统日志", map[string]interface{}{
		"content":     fileContent,   // 文件内容
		"current":     postData.Name, // 文件名称
		"currentLine": fileLine,      // 开始显示的文件行数内容(0为第一行)
		"lists":       logLists,
		"size":        fileSize, // 字节
		"types":       showDir,
	})
	return
}

// DeleteSystemLogs 删除系统日志文件
func (*System) DeleteSystemLogs(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var (
		PostData manageBusiness.DeleteSystemLogsSt
		err      error
	)

	PostData.Name, err = zvalid.Text(c.GetJSON("name").Str, "文件名称").Required().String()
	if err != nil {
		global.Log.Error("参数错误: ", PostData)
		c.ApiJSON(200, "删除系统日志", false)
		return
	}
	PostData.Type, err = zvalid.Text(c.GetJSON("type").Str, "文件夹名称").Required().String()
	if err != nil {
		global.Log.Error("参数错误: ", PostData)
		c.ApiJSON(200, "删除系统日志", false)
		return
	}

	rmLogPath := zfile.RealPath("./" + PostData.Type + "/" + PostData.Name)
	// 因为当天日志一直被占用 无法删除
	// 又因为日志生成都是日期组成所以..
	nowDate := time.Now().Format("2006-01-02")
	if PostData.Name == nowDate {
		global.Log.Error("删除文件得时间和参数输入的时间相等,所以直接清空内容.")
		f, _ := os.Create(rmLogPath)
		_ = f.Close()
		c.ApiJSON(200, "删除系统日志", true)
		return
	}

	_, err = os.Lstat(rmLogPath)
	if err == nil {
		err = os.Remove(rmLogPath)
		if err != nil {
			global.Log.Error("删除文件失败: ", err)
			c.ApiJSON(200, "删除系统日志", false)
			return
		}

		global.Log.Error("删除文件成功")
		c.ApiJSON(200, "删除系统日志", true)
		return
	}

	c.ApiJSON(200, "删除系统日志", false)
	return
}

// GetSystemConfig 读取系统配置
func (*System) GetSystemConfig(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var paramPutSystemConfigSt manageBusiness.ParamPutSystemConfigSt
	res, err := paramPutSystemConfigSt.GetConf()
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "读取系统配置", res)
}

// PutSystemConfig 更新系统配置
func (*System) PutSystemConfig(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var paramPutSystemConfigSt manageBusiness.ParamPutSystemConfigSt
	tempRule := c.ValidRule()
	err := zvalid.Batch(
		zvalid.BatchVar(&paramPutSystemConfigSt.IpWhitelist, c.Valid(tempRule, "ipWhitelist", "IP白名单")),
		zvalid.BatchVar(&paramPutSystemConfigSt.MaintainMode, c.Valid(tempRule, "maintainMode", "维护模式")),
		zvalid.BatchVar(&paramPutSystemConfigSt.Debug, c.Valid(tempRule, "debug", "开发模式")),
		zvalid.BatchVar(&paramPutSystemConfigSt.CdnHost, c.Valid(tempRule, "cdnHost", "cdn地址")),
		zvalid.BatchVar(&paramPutSystemConfigSt.LoginMode, c.Valid(tempRule, "loginMode", "登录模式")),
	)
	if err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}
	if err := paramPutSystemConfigSt.SetConf(); err != nil {
		c.ApiJSON(211, err.Error(), nil)
		return
	}

	if paramPutSystemConfigSt.LoginMode {
		t, _ := c.Value("token")
		_ = t.(*model.AuthUserToken).LoginModeTrue()
	}

	c.ApiJSON(200, "更新系统配置", true)
}

// GetMenu 获取系统菜单列表
func (*System) GetMenu(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}
}
