package manage

import (
	"app/model"
	"app/web/business/manageBusiness"
	"fmt"
	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/znet"
	"github.com/sohaha/zlsgo/zvalid"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

// 后台-系统接口
type System struct {
}

// GetLogs 查看用户日志
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

// GetUnreadMessageCount 未读日志总数
func (*System) GetUnreadMessageCount(c *znet.Context) {
	lastId, _ := strconv.Atoi(c.DefaultFormOrQuery("id", "0"))

	u, _ := c.Value("user")
	userid := u.(*model.AuthUser).ID

	c.ApiJSON(200, "未读日志", (&model.AuthUserLogs{Userid: userid, ID: uint(lastId)}).UnreadMessageCount())
	return
}

// PutMessageStatus 更新日志状态
func (*System) PutMessageStatus(c *znet.Context) {
	idsMap, _ := c.GetPostFormMap("ids")

	u, _ := c.Value("user")
	uid := u.(*model.AuthUser).ID

	ids := []int{}
	for _, v := range idsMap {
		i, _ := strconv.Atoi(v)
		ids = append(ids, i)
	}

	count := (&model.AuthUserLogs{Userid: uid}).UpdateMessageStatus(ids)
	c.ApiJSON(200, "日志标记已读", count)
	return
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

// DeleteSystemLogs 删除系统日志文件
func (*System) DeleteSystemLogs(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}

	var PostData manageBusiness.DeleteSystemLogsSt

	temRule := c.ValidRule().Required()
	err := zvalid.Batch(
		zvalid.BatchVar(&PostData.Name, c.Valid(temRule, "name", "文件名称")),
		zvalid.BatchVar(&PostData.Type, c.Valid(temRule, "type", "文件夹名称")),
	)

	if err != nil {
		zlog.Error("参数错误: ", PostData)
		c.ApiJSON(200, "删除系统日志", false)
		return
	}

	rmLogPath := zfile.RealPath("./" + PostData.Type + "/" + PostData.Name)
	// 因为当天日志一直被占用 无法删除
	// 又因为日志生成都是日期组成所以..
	nowDate := time.Now().Format("2006-01-02")
	zlog.Error("当前时间: ", nowDate)
	if PostData.Name == nowDate {
		zlog.Error("删除文件得时间和参数输入的时间相等,所以直接清空内容.")
		f, _ := os.Create(rmLogPath)
		f.Close()
		c.ApiJSON(200, "删除系统日志", true)
		return
	}

	_, err = os.Lstat(rmLogPath)
	zlog.Error("文件是否存在: ", err)
	if err == nil {
		// f, _ := os.Open(rmLogPath)
		// f.Close()
		err = os.Remove(rmLogPath)
		if err != nil {
			zlog.Error("删除文件失败: ", err)
			c.ApiJSON(200, "删除系统日志", false)
			return
		}

		zlog.Error("删除文件成功")
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
		c.ApiJSON(201, err.Error(), nil)
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

// GetMenu 获取系统菜单列表
func (*System) GetMenu(c *znet.Context) {
	if !VerifPermissionMark(c, "systems") {
		return
	}
}
