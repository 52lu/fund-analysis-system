package initialize

import (
	"52lu/fund-analye-system/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path"
	"runtime"
)

// 获取配置文件位置
func getConfigFile() string {
	_, file, _, _ := runtime.Caller(1)
	// 拼凑配置文件路径
	return path.Dir(file)+"/../config.yaml"
}
// ViperInit 初始化viper配置解析包，函数可接受命令行参数
func initConfig() {
	var configFile string
	// 读取配置文件优先级: 命令行 > 默认值
	_, file, _, _ := runtime.Caller(1)
	fmt.Println("path.file = ", file)
	fmt.Println("path.Dir = ", path.Dir(file))

	flag.StringVar(&configFile,"c",getConfigFile(),"配置配置")
	if len(configFile) == 0 {
		// 读取默认配置文件
		panic("配置文件不存在！")
	}
	// 读取配置文件
	v := viper.New()
    v.SetConfigFile(configFile)
	if err := v.ReadInConfig();err != nil {
		panic(fmt.Errorf("配置解析失败:%s\n",err))
	}
	// 动态监测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生改变")
		if err := v.Unmarshal(&global.GvaConfig); err != nil {
			panic(fmt.Errorf("配置重载失败:%s\n",err))
		}
	})
	if err := v.Unmarshal(&global.GvaConfig); err != nil {
		panic(fmt.Errorf("配置重载失败:%s\n",err))
	}
	// 设置配置文件
	global.GvaConfig.App.ConfigFile = configFile
}
