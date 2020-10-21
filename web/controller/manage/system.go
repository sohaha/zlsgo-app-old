package manage

import (
	"app/model"
	"app/web/business/manageBusiness"
	"fmt"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"
	"io/ioutil"
	"strconv"
	"strings"
)

type System struct {
}

// 查看用户日志
func (*System) GetLogs(c *znet.Context) {
	pagesize, _ := strconv.Atoi(c.DefaultFormOrQuery("pagesize", "10"))
	page, _ := strconv.Atoi(c.DefaultFormOrQuery("page", "1"))
	qType, _ := strconv.Atoi(c.DefaultFormOrQuery("type", "0"))
	qUnread, _ := strconv.Atoi(c.DefaultFormOrQuery("unread", "0"))

	u, _ := c.Value("user")
	userid := u.(*model.AuthUser).ID

	p := model.Page{
		Curpage:  uint(page),
		Pagesize: uint(pagesize),
	}

	logs := (&model.AuthUserLogs{Userid: userid, Type: uint8(qType), Status: uint8(qUnread)}).Lists(&p)

	c.ApiJSON(200, "用户日志", map[string]interface{}{
		"items": logs,
		"page":  p,
	})
}

// 未读日志总数
func (*System) GetUnreadMessageCount(c *znet.Context) {}

// 更新日志状态
func (*System) PutMessageStatus(c *znet.Context) {}

// 系统日志
func (*System) GetSystemLogs(c *znet.Context) {
	var postData manageBusiness.ParamGetSystemLogsSt
	tempRule := c.ValidRule()
	err := zvalid.Batch(
		zvalid.BatchVar(&postData.Name, c.Valid(tempRule, "name", "文件名称")),
		zvalid.BatchVar(&postData.Type, c.Valid(tempRule, "type", "文件夹名称")),
		zvalid.BatchVar(&postData.CurrentLine, c.Valid(tempRule, "currentLine", "当前行")),
	)
	if err != nil {
		c.ApiJSON(201, err.Error(), nil)
		return
	}

	showDir := manageBusiness.GetTmpLogDir()
	logLists := manageBusiness.ShowLogsLists(postData.Type)
	inDir := false
	for _, dir := range showDir {
		if dir == postData.Type {
			inDir = true
		}
	}

	logItems := make([]string, 0)
	if postData.Name != "" && inDir {
		// aaa := zfile.RealPath("./static/logs/"+time.Now().Format("2006"))
		searchDirPath := zfile.RealPath("./" + postData.Name)

		dirInfo, _ := ioutil.ReadDir(searchDirPath)
		for _, v := range dirInfo {
			if v.IsDir() {
				lastDir := searchDirPath + "/" + v.Name()
				lastDirInfo, _ := ioutil.ReadDir(lastDir)
				for _, last := range lastDirInfo {
					if !last.IsDir() {
						logItems = append(logItems, v.Name()+"/"+last.Name())
					}
				}
			}
		}
	}

	logPath := zfile.RealPath("./" + postData.Type + "/" + postData.Name)

	fileContentSlice, fileSize, fileLine := manageBusiness.GetSystemLogInfo(logPath, postData.CurrentLine)
	fileContent := ""
	if fileSize > (1024 * 500) {
		fileContent = "日志文件过大（" + fmt.Sprintf("%v", manageBusiness.Byte2Kb(fileSize)) + "kb），不支持在线查看全部内容。\n\n"
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

// 删除系统日志文件
func (*System) DeleteSystemLogs(c *znet.Context) {}

// 读取系统配置
func (*System) GetSystemConfig(c *znet.Context) {
	var paramPutSystemConfigSt manageBusiness.ParamPutSystemConfigSt
	res, err := paramPutSystemConfigSt.GetConf()
	if err != nil {
		c.ApiJSON(201, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "读取系统配置", res)
}

// 更新系统配置
func (*System) PutSystemConfig(c *znet.Context) {
	var paramPutSystemConfigSt manageBusiness.ParamPutSystemConfigSt
	tempRule := c.ValidRule()
	err := zvalid.Batch(
		zvalid.BatchVar(&paramPutSystemConfigSt.IpWhitelist, c.Valid(tempRule, "ipWhitelist", "IP白名单")),
		zvalid.BatchVar(&paramPutSystemConfigSt.MaintainMode, c.Valid(tempRule, "maintainMode", "维护模式")),
		zvalid.BatchVar(&paramPutSystemConfigSt.Debug, c.Valid(tempRule, "debug", "开发模式")),
		zvalid.BatchVar(&paramPutSystemConfigSt.CdnHost, c.Valid(tempRule, "cdnHost", "cdn地址")),
	)
	if err != nil {
		c.ApiJSON(201, err.Error(), nil)
		return
	}
	if err := paramPutSystemConfigSt.SetConf(); err != nil {
		c.ApiJSON(201, err.Error(), nil)
		return
	}

	c.ApiJSON(200, "更新系统配置", true)
}

// 获取系统菜单列表
func (*System) GetMenu(c *znet.Context) {}
