package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	dingdingToken string //钉钉告警url
	podId         string //PodId
	folder        string //日期命名文件夹
	postfix       string //文件名的时间后缀
	env           string //部署环境
	projectGroup  string //项目组
	projectArray  string //PodId切片后的数组
	projectName   string //项目名称
	bucketName    string //OSS bucketName
	locaFilename  string //OOM DumpFile
)

func init() {

	flag.StringVar(&podId, "k", "ops", "PodId")
	flag.StringVar(&env, "e", "test", "ENV")

	folder = time.Now().Format("20060102")
	postfix = time.Now().Format("20060102150405")
	locaFilename = fmt.Sprint("log/java_heapdump.hprof")
}

// 判断所给路径文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	// 如果文件存在，返回true
	if err == nil {
		return true, nil
	}
	// 如果文件不存在，返回false
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 获取projectName
func GetProjectName() string {
	projectArray := strings.Split(podId, "-")
	return fmt.Sprintf("%s-%s-%s", projectArray[0], projectArray[1], projectArray[2])
}

// 删除原有dump文件
func RemoveDumpfile(path string) {
	os.Remove(path)

	exist, err := PathExists(locaFilename)

	if exist {
		fmt.Printf("panic %s\n", err)
		return
	} else {
		fmt.Printf("remove %s completely\n", path)
		return
	}
}

func main() {
	flag.Parse()

	// 异常处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic %s\n", err)
		}
	}()

	projectName = GetProjectName()
	print(fmt.Sprintf("ENV: %s PROJECT: %s POD: %s JVM OOM occurs!", env, projectName, podId))
	// projectGroup = fmt.Sprintf(strings.Split(podId, "-")[0]) // podId: "ops-demo"
	// bucketName = fmt.Sprintf("%s-disaster", projectGroup)    //正式的bucketName，如ops-disaster

	projectGroup = "nanjing-java"
	bucketName = "uniondrug-k8s"

	// 判断dump文件是否存在
	exist, err := PathExists(locaFilename)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		alarm()
		upload()
		RemoveDumpfile(locaFilename)
	}
}
