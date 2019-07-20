# adby
一个基于adb的包管理器，类似于Fedora的dnf/yum，以及Ubuntu的apt。


## 支持的子命令
目前支持的子命令有：

- clean
    - cache 清理缓存的应用安装包
- download <包名>，下载应用
- install <包名>，安装应用
    - -l，从本地进行安装
- uninstall <包名>，卸载应用
- info <包名>，获取应用的基本信息
- list，列出应用信息，默认列出全部已安装的应用
  - -s，列出系统预置应用
  - -3，列出第三方应用
- search <名称关键字>，应用搜索
- [TODO] update <包名列表>，更新应用
- [TODO] upgrade <包名列表，默认位全部应用>，更新应用

- help，获取应用的帮助信息
- version，adby版本号
