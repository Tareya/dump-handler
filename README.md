
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
- OOM Dump文件路径
  - /data/apps/${projectName}/log/java_heapdump.hprof


### 参考文档
[处理k8s中java应用OOM时的dump文件(非preStop)](http://www.devopser.org/articles/2020/09/17/1600339403553.html)
