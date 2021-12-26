package conf

type (
	WebConf struct {
		Port       string // 项目端口
		Tls        bool   // 开启 https
		TlsPort    string // https 端口
		Key        string // 证书
		Cert       string // 证书
		Debug      bool   // 开启调试模式
		Pprof      bool   // 开启 pprof
		PprofToken string // pprof Token
	}
)

func init() {
	allConf["web"] = &WebConf{
		Port: "3788",
	}
}

// Web Web配置
//goland:noinspection ALL
func Web() WebConf {
	conf := getConf(func() interface{} {
		conf, ok := allConf["web"]
		if !ok {
			return WebConf{}
		}
		return *conf.(*WebConf)
	})

	return conf.(WebConf)
}
