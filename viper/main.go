package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
	//_ "github.com/spf13/viper/remote"
)

var (
	cfg   = pflag.StringP("config", "c", "", "Configuration file.")
	token = pflag.String("token", "", "Token to access http service.")
	help  = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		viper.AddConfigPath(".")          // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("$HOME/.iam") // 配置文件搜索路径，可以设置多个配置文件搜索路径
		viper.SetConfigName("config")     // 配置文件名称（没有文件扩展名）
	}

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 打印当前使用的配置文件名
	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())

	// 修改配置
	viper.Set("user.username", "colin")
	viper.SetDefault("max-retries", 4)

	// 注册和使用别名
	viper.RegisterAlias("mr", "max-retries")

	// 使用环境变量
	os.Setenv("VIPER_USER_SECRET_ID", "QLdywI2MrmDVjSSv6e95weNRvmteRjfKAuNV")
	os.Setenv("VIPER_USER_SECRET_KEY", "bVix2WBv0VPfrDrvlLWrhEdzjLpPCNYb")

	viper.AutomaticEnv()                                             // 读取环境变量
	viper.SetEnvPrefix("VIPER")                                      // 设置环境变量前缀：VIPER_，如果是viper，将自动转变为大写。
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_")) // 将viper.Get(key) key字符串中'.'和'-'替换为'_'
	viper.BindEnv("user.secret-key")
	viper.BindEnv("user.secret-id", "USER_SECRET_ID") // 绑定环境变量名到key

	// 绑定标志
	viper.BindPFlag("token", pflag.Lookup("token")) // 绑定单个标志
	viper.BindPFlags(pflag.CommandLine)             //绑定标志集

	// 远程Key/Value存储示例-未加密（etcd）
	/*
		viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yaml")
		viper.SetConfigType("yaml") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。
		if err := viper.ReadRemoteConfig(); err != nil {
			panic(fmt.Errorf("Read configuration from etcd: %v \n", err))
		}
	*/

	// 监控etcd中的更改-未加密
	/*
		// 或者你可以创建一个新的viper实例
		var runtime_viper = viper.New()

		runtime_viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001", "/config/hugo.yml")
		runtime_viper.SetConfigType("yaml") // 因为在字节流中没有文件扩展名，所以这里需要设置下类型。

		// 第一次从远程读取配置
		err := runtime_viper.ReadRemoteConfig()

		// 反序列化
		runtime_viper.Unmarshal(&runtime_conf)

		// 开启一个单独的goroutine一直监控远端的变更
		go func() {
			for {
				time.Sleep(time.Second * 5) // 每次请求后延迟一下

				// 目前只测试了etcd支持
				err := runtime_viper.WatchRemoteConfig()
				if err != nil {
					log.Errorf("unable to read remote config: %v", err)
					continue
				}

				// 将新配置反序列化到我们运行时的配置结构体中。你还可以借助channel实现一个通知系统更改的信号
				runtime_viper.Unmarshal(&runtime_conf)
			}
		}()
	*/

	// 打印配置
	fmt.Println("apiVersion = ", viper.GetString("apiVersion"))
	fmt.Println("max-retries = ", viper.GetString("max-retries"))
	fmt.Println("mr = ", viper.GetString("mr"))
	fmt.Println("user.username = ", viper.GetString("user.username")) // 访问嵌套的键
	fmt.Println("user.password = ", viper.GetString("user.password"))
	fmt.Println("user.secret-id = ", viper.GetString("user.secret-id"))
	fmt.Println("user.secret-key = ", viper.GetString("user.secret-key"))
	fmt.Println("token = ", viper.GetString("token"))
	fmt.Println("config = ", viper.GetString("config"))

	// 反序列化
	var config struct {
		APIVersion string `mapstructure:"apiVersion"`
		MaxRetries int    `mapstructure:"max-retries"`
		User       struct {
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
		} `mapstructure:"user"`
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	fmt.Println("apiVersion(struct) = ", config.APIVersion)
	fmt.Println("user.username(struct) = ", config.User.Username)

	// 序列化成字符串
	configString := yamlStringSettings()
	fmt.Printf("\nmarshal config to YAML:\n%s", configString)

	// 将viper当前的配置写入配置文件
	//viper.WriteConfig()                            // 保存当前的配置到viper当前使用的配置文件中，如果配置文件不存在会报错，如果配置文件存在则覆盖当前的配置文件
	//viper.SafeWriteConfig()                        // 保存当前的配置到viper当前使用的配置文件中，如果配置文件不存在会报错，如果配置文件存在则返回file exists错误
	viper.WriteConfigAs("config.running.yaml") // 保存当前的配置到指定的文件中，如果文件不存在则新建，如果文件存在则会覆盖文件
	//viper.SafeWriteConfigAs("config.running.yaml") // 保存当前的配置到指定的文件中，如果文件不存在则新建，如果文件存在则返回file exists错误

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}

func yamlStringSettings() string {
	c := viper.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		panic(fmt.Errorf("unable to marshal config to YAML: %v", err))
	}
	return string(bs)
}
