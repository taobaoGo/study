# 安装&环境
### wsl ssh
```shell
cp -rf /mnt/c/Users/admin/.ssh /root
## Permissions 0777 for '/root/.ssh/id_rsa' are too open.
chmod 700 /root/.ssh
```
### wsl go env
```shell
# https://golang.google.cn/dl/
wget https://dl.google.com/go/go1.12.7.linux-amd64.tar.gz
mkdir -p /usr/local
tar -C /usr/local -xzf go1.12.7.linux-amd64.tar.gz
sudo vim /etc/profile
## 在文件末尾追加
export GOROOT=/usr/local/go
export GOPATH=/mnt/d/gopath
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
# 保存完成后执行
source /etc/profile
```
## 工作区work space
### src：源码存放目录
> 存放*.go文件。
- 代码包:src/projname/下的文件夹，源码文件的包名与文件夹同名。
### pkg：归档文件目录
>
存放*.a文件。归档文件名就是源码文件的文件夹名。
### bin：存放可执行文件
>
文件名是项目名，即命令源码文件所在的文件夹名。

## 包管理
https://www.jianshu.com/p/a537ee63d606
### go modules
https://github.com/golang/go/wiki/Modules
```shell
go mod init ProjectName
go build -o ProjectName
```

# IDE:Goland
## 数据库工具
Goland-->View-->Tools Window-->Database


 

