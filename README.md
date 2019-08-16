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
# IDE
## Goland中的数据库工具
Goland-->View-->Tools Window-->Database

### package
https://www.jianshu.com/p/a537ee63d606

# 工作区work space
## mod

## src：源码存放目录，存放*.go文件。
- 代码包:就是src目录下的文件夹，一个文件夹下的源代码必须只有一个包名。
## pkg：归档文件目录，存放*.a文件。
## bin：存放可执行文件。
