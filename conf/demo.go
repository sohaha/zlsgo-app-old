package conf

type (
	DemoConf struct {
		Test string
	}
)

func init() {
	allConf["demo"] = &DemoConf{
		// Test: "test", // 默认配置
	}
}

// Demo Demo配置
//goland:noinspection ALL
func Demo() DemoConf {
	conf := getConf(func() interface{} {
		conf, ok := allConf["demo"]
		if !ok {
			return DemoConf{}
		}
		return *conf.(*DemoConf)
	})

	return conf.(DemoConf)
}
