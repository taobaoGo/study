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

## Goland中的数据库工具
Goland-->View-->Tools Window-->Database

### package
https://www.jianshu.com/p/a537ee63d606
