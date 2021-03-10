# cobra课后练习

1. 准备，参考：[实战练习准备](../../docs/prepare.md)


2. 安装cobra命令

```bash
$ go get -u github.com/spf13/cobra/cobra
```

3. 创建demo

```bash
$ mkdir -p cobra
$ cd cobra
$ cobra init --viper --pkg-name='github.com/marmotedu/gopractise-demo/cobra/demoapp' demoapp
```

3. 创建命令

```bash
$ cd demoapp
$ cobra add config
$ cobra add serve
```

4. 编译运行

```bash
$ go build -v .
$ ./demoapp config
$ ./demoapp serve
```




