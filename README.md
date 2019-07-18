# adby
一个基于adb的包管理器，类似于Fedora的dnf/yum，以及Ubuntu的apt。


## 支持的子命令
初步计划支持的子命令有：

- download <包名>，下载应用
- install <包名>，安装应用
- uninstall <包名>，卸载应用
- info <包名>，获取应用的基本信息
  - -p，接包名
  - -n，接应用名称
- list，列出应用信息，默认列出全部已安装的应用
  - -s，列出系统预置应用
  - -3，列出第三方应用
- search <名称关键字>，应用搜索
- update <包名列表>，更新应用
- upgrade <包名列表，默认位全部应用>，更新应用

- help，获取应用的帮助信息
- version，adby版本号
