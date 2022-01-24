## 简介
viper 是一个配置解决方案，拥有丰富的特性：

支持 JSON/TOML/YAML/HCL/envfile/Java properties 等多种格式的配置文件；（在internal/encoding下）
可以设置监听配置文件的修改，修改时自动加载新的配置；
从环境变量、命令行选项和io.Reader中读取配置；
从远程配置系统中读取和监听修改，如 etcd/Consul；
代码逻辑中显示设置键值。

## 仓库
repo https://github.com/spf13/viper
## 读取配置
``` go
viper.Set("test", "test")  // 会覆盖对应读取的配置
```

## 写入配置

## 监听文件修改

