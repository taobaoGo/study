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
export GOROOT=/usr/local/go
export GOPATH=/mnt/d/gopath
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```
