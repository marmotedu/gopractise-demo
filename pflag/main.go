package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

const helpText = `Usage: main [flags] arg [arg...]

This is a pflag example.

Flags:
`

// 定义命令行参数，并将标志的值存储在指针中.
var (
	username           = pflag.String("username", "root", "Username for access to mysql service.")
	password           = pflag.StringP("password", "p", "root", "Password for access to mysql, should be used pair with password.")
	maxIdleConnections = pflag.Int("max-idle-connections", 100, "Maximum idle connections allowed to connect to mysql.")
	timeout            = pflag.DurationP("timeout", "t", time.Duration(10)*time.Second, "Timeout when connecting to mysql service.")
	logmode            = pflag.BoolP("logmode", "l", false, "Start gorm debug mode.")
	logMode            = pflag.BoolP("log-mode", "m", false, "Start gorm debug mode.")
	help               = pflag.BoolP("help", "h", false, "Show this help message.")

	usage = func() {
		fmt.Println(helpText)
		pflag.PrintDefaults()
	}
)

func wordSepNomalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}

func main() {
	var (
		host string
		port int
	)

	// 定义命令行参数，并将标志的值绑定到变量
	pflag.StringVar(&host, "host", "127.0.0.1", "MySQL service host address.")
	pflag.IntVarP(&port, "port", "P", 3306, "MySQL service host port.")

	// 设置标准化参数名称的函数
	pflag.CommandLine.SetNormalizeFunc(wordSepNomalizeFunc)

	// 为 username 参数设置 NoOptDefVal
	pflag.Lookup("username").NoOptDefVal = "colin"

	// 把 logmode 参数标记为即将废弃的，请用户使用 --log-mode
	_ = pflag.CommandLine.MarkDeprecated("logmode", "please use --log-mode instead")

	// 把 port 参数的 shorthand 标记为即将废弃的，请用户使用 --port
	_ = pflag.CommandLine.MarkShorthandDeprecated("port", "please use --port only")

	// 在帮助文档中隐藏参数 max-idle-connections
	_ = pflag.CommandLine.MarkHidden("max-idle-connections")

	pflag.Usage = usage

	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()

	if *help {
		pflag.Usage()
		return
	}

	// 获取命令行参数的值
	// 获取选项参数的值
	fmt.Println("username =", *username)
	fmt.Println("password =", *password)
	fmt.Println("max-idle-connections =", *maxIdleConnections)
	fmt.Println("timeout =", *timeout)
	fmt.Println("logmode =", *logmode)
	fmt.Println("log-mode =", *logMode)
	fmt.Println("host =", host)

	p, _ := pflag.CommandLine.GetInt("port")
	fmt.Println("port =", p)

	// 获取非选项参数的值
	fmt.Println("\nArgs = ", pflag.Args())
	fmt.Println("Args[0] = ", pflag.Arg(0))

	// 获取非选项参数和Flag的数量
	fmt.Println("Number of Args = ", pflag.NArg())
	fmt.Println("Number of Flags = ", pflag.NFlag())
}
