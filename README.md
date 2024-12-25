 pssh

## 说明
这是一个使用golang写的命令行工具，用来使用ssh命令通过密码的形式登录主机的。
    
## 开发初衷
本人使用的mac系统，使用的iTerm2终端，做运维工作需要经常登录linux服务器，大部分服务器都使用密钥登录，或者是堡垒机来管理，但个别服务器使用都要输入密码，非常麻烦，而iterm2又不像win下的Xshell等管理工具提供密码登录功能，所以写了这个工具。
    本人也是go的初学者，望大佬多提意见，欢迎指正。
## 依赖
1. 本程序在实现ssh登录时使用了`golang.org/x/crypto/ssh`, 需要`go`版本在`1.20`以上,其他版本没有测试过。

2. 本程序使用了`sqlite3`做为存储数据库.

# 构建
```bash
go build -ldflags="-w -s" -o pssh main.go
```

## 使用
```bash
pssh -h 
  add         Add host to database.
  del         Delete a host from the database
  help        Help about any command
  list        List all the hosts in the database
  login       Login to host
  init        Initialize the database
  version     Print the version number of pssh
```
#### 命令说明
- --init: 初始化数据库，创建存储主机的表`hosts`
- --add: 添加主机到数据库，参数依次为`name、ip、port、user、password`
- --del : 删除主机,可以是主机名，也可以是ip
- list: 显示数据库中所有主机
- login: 登录主机
#### 示例
```bash
# 添加主机，参数为名称、ip、端口、用户名、密码
pssh add testhost1 192.168.1.100 22 root 123456
# 删除主机, 参数可以为主机名，也可以为ip
pssh del testhost1
pssh del 192.168.1.100
# 显示数据库中所有主机, -a显示所有信息
pssh list [--all|-a]
# 登录主机, 参数可以为主机名，也可以为ip
pssh login testhost1
pssh login 192.168.1.100
```
### 相关配置
1. 数据库存放位置：`$HOME/.pssh/hosts.db`, 可以在文件`database/init.go`修改`dbpath`。
2. 主机的密码是经过对称加密后存储到数据库的，加密的密钥在`utils/utils.go`里的`key`, 默认`key`的值为`LinuxKaiLinuxKai`