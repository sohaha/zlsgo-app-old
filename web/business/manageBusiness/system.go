package manageBusiness

import (
	"bufio"
	"github.com/sohaha/gconf"
	"github.com/sohaha/zlsgo/zfile"
	"io"
	"io/ioutil"
	"os"
)

type ParamGetSystemLogsSt struct {
	Name        string
	Type        string
	CurrentLine int
}

// 返回log目录列表
func GetTmpLogDir() []string {
	path := zfile.RealPath(".")
	poolDir := map[string]int8{
		"logs": 1,
	}
	reDir := make([]string, 0)
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			_, has := poolDir[file.Name()]
			if has {
				reDir = append(reDir, file.Name())
			}
		}
	}

	return reDir
}

func ShowLogsLists(logType string) []string {
	reLists := make([]string, 0)

	// 这里会默认指向运行时候设置的,例如tmp

	logPath := zfile.RealPath(".")
	showDir := GetTmpLogDir()
	if logType != "" {
		for _, dir := range showDir {
			if dir != logType {
				continue
			}

			temDirPath := logPath + "/" + dir
			dirInfo, _ := ioutil.ReadDir(temDirPath)

			for _, info := range dirInfo {
				if info.IsDir() {
					continue
				}

				reLists = append(reLists, info.Name())
			}
		}
	} else {
		for _, dir := range showDir {
			temDirPath := logPath + "/" + dir
			dirInfo, _ := ioutil.ReadDir(temDirPath)
			for _, info := range dirInfo {
				if info.IsDir() {
					continue
				}
				reLists = append(reLists, dir+"/"+info.Name())
			}
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

	// currentCfg := cfg.GetAll()
	st.Debug = cfg.Get("base.debug").(bool)
	st.CdnHost = cfg.Get("project.cdnHost").(string)
	st.MaintainMode = cfg.Get("base.maintainMode").(bool)
	st.IpWhitelist = cfg.Get("base.ipWhitelist").(string)
	return st, nil
}
