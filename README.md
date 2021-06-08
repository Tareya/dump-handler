
### 功能：

判断是否存在jvm dump文件"java_heapdump.hprof"，如果存在就把dump文件上传至oss，并发送钉钉告警，如果不存在则忽略。
oss文件分片，断点续传
钉钉告警，附带oss链接地址，根据podid自动判断发送到对应项目组的告警群

### 编译：

```
1. mac 可执行
export GO111MODULE=on && export GOPROXY="https://goproxy.cn,direct" && CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o build/`date "+%Y%m%d%H%M"`/dump-handler

2. linux 可执行
export GO111MODULE=on && export GOPROXY="https://goproxy.cn,direct" && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/`date "+%Y%m%d%H%M"`/dump-handler

3. windows 可执行
export GO111MODULE=on && export GOPROXY="https://goproxy.cn,direct" && CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-w -s" -o build/`date "+%Y%m%d%H%M"`/dump-handler
```

### 使用方法：

```
1、添加jvm参数，当应用发生OOM时会自动执行工具"-XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=${BASE_DIR}/log/java_heapdump.hprof -XX:+ExitOnOutOfMemoryError -XX:OnOutOfMemoryError=./dump-handler -k \$HOSTNAME -e \$ENV"
2、部署应用到k8s时在deployment配置挂载日志目录 /mnt/${projectName}/log
```

### 说明：

- PODID
  - k8s pod的id，可以通过$HOSTNAME获取
  - podid命名规范，示例 "uniondrug-pc-web-6cd649945c-8ztmp"，以"-"为分隔符
- ENV
  - 部署环境，可以通过$ENV获取(提前通过deployment配置环境变量ENV到容器中)
- OOM Dump文件路径(相对路径，要求编译文件在构建docker镜像时，直接上传至workspace中)
  - log/java_heapdump.hprof


### 可更改变量以及作用

#### main.go
- projectGroup
  - 项目分组，可以在有多个项目组拆分的时候定义项目组

- bucketName
  - oss bucket，用于需要拆分bucket的情况

- locaFilename
  - dump文件的路径


#### oss.go
- endpoint
  - oss endpoint url，正式环境推荐是用 internal

- accessKeyID
- accessKeySecret
  - aliyun 的 AK

- objectName
  - 上传 oss 的文件路径

#### dingtalk.go
- tokenMap
  - dingding机器人的 token 集合，根据 projectGroup 做map

- ossUrl
  - 访问oss存储文件的路径（使用公网endpoint以便下载）

- alarmMsg
  - 发送到dingding机器人的消息体，markdown格式