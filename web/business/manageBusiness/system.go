package manageBusiness

import (
	"bufio"
	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zfile"
	"io"
	"io/ioutil"
	"os"
)

type (
	GetSystemLogsSt struct {
		Name        string
		Type        string
		CurrentLine int
	}

	DeleteSystemLogsSt struct {
		Name        string
		Type        string
		CurrentLine int
	}
)

// 返回log目录列表
func GetTmpLogDir(logDir string) []string {
	var reDir []string
	if logDir == "" {
		return reDir
	}

	path := zfile.RealPath("./" + logDir)
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			reDir = append(reDir, file.Name())
		}
	}

	return reDir
}

func ShowLogsLists(logType string, logDir string) []string {
	var reLists []string
	if logDir == "" {
		return reLists
	}

	logPath := zfile.RealPath("./" + logDir)
	dirInfo, _ := ioutil.ReadDir(logPath + "/" + logType)
	for _, info := range dirInfo {
		if !info.IsDir() {
			reLists = append(reLists, info.Name())
		}
	}

	return reLists
}

// filePath 文件路径
// readLine 第几行开始输出
func GetSystemLogInfo(filePath string, readLine int) (fileContent []string, fileSize int64, fileLine int) {
	fi, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer fi.Close()

	// syscall\syscall_windows.go
	// func Seek
	// 2 末尾
	// 0 开始
	// os.SEEK_END
	// fi.Seek(0, 2)
	fileSize, _ = fi.Seek(0, 2)
	fi.Seek(0, 0)

	buf := bufio.NewReader(fi)
	for {
		data, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			break
		}

		if fileLine >= readLine {
			fileContent = append(fileContent, data)
		}

		fileLine++
	}

	// fmt.Println(content[len(content)-10:])
	// fmt.Println(len(content))

	// fileContent = strings.Join(content, "")
	// fmt.Println("文件内容: ", fileContent)
	// fmt.Println("容量大小: ", adminLogic.Byte2Kb(fileSize))
	// fmt.Println("行数: ", fileLine)

	return
}

type ParamPutSystemConfigSt struct {
	IpWhitelist  string `json:"ipWhitelist"`
	MaintainMode bool   `json:"maintainMode"`
	Debug        bool   `json:"debug"`
	CdnHost      string `json:"cdnHost"`
	LoginMode    bool   `json:"loginMode"`
}

func (st *ParamPutSystemConfigSt) SetConf() error {
	Mutex.Lock()

	cfgName := "conf.yml"
	cfg := gconf.New(cfgName)
	err := cfg.Read()
	if err != nil {
		return err
	}

	cfg.Set("base.debug", st.Debug)
	cfg.Set("project.cdnHost", st.CdnHost)
	cfg.Set("base.maintainMode", st.MaintainMode)
	cfg.Set("base.ipWhitelist", st.IpWhitelist)
	cfg.Set("base.loginMode", st.LoginMode)
	cfg.Write(cfgName)

	Mutex.Unlock()
	return nil
}

func (st *ParamPutSystemConfigSt) GetConf() (*ParamPutSystemConfigSt, error) {
	cfgName := "conf.yml"
	cfg := gconf.New(cfgName)
	err := cfg.Read()
	if err != nil {
		return nil, err
	}

	cfgDebug := cfg.Get("base.debug")
	if cfgDebug != nil {
		st.Debug = cfgDebug.(bool)
	}
	cfgCdnHost := cfg.Get("project.cdnHost")
	if cfgCdnHost != nil {
		st.CdnHost = cfgCdnHost.(string)
	}
	cfgMaintainMode := cfg.Get("base.maintainMode")
	if cfgMaintainMode != nil {
		st.MaintainMode = cfgMaintainMode.(bool)
	}
	cfgIpWhitelist := cfg.Get("base.ipWhitelist")
	if cfgIpWhitelist != nil {
		st.IpWhitelist = cfgIpWhitelist.(string)
	}
	cfgLoginMode := cfg.Get("base.loginMode")
	if cfgLoginMode != nil {
		st.LoginMode = cfgLoginMode.(bool)
	}

	return st, nil
}
