sudo vim /etc/profile
## 在文件末尾追加
export GOROOT=/usr/local/go
export GOPATH=/mnt/d/gopath
export GOBIN=$GOPATH/bin
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
# 保存完成后执行
source /etc/profile