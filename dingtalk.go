package main

import (
	"fmt"
	"time"

	"github.com/braumye/grobot"
)

func msgSender(dingdingToken string, msg string) {
	robot, _ := grobot.New("dingtalk", dingdingToken)
	err := robot.SendMarkdownMessage("通知", msg)
	fmt.Println("通知发送完毕", err)
}

func alarm() {
	//key=项目组，value=钉钉token [需要修改]
	tokenMap := map[string]string{
		"nanjing-java": "https://oapi.dingtalk.com/robot/send?access_token=dde2598d393894a05baddf95216ea4796785b2a011290017419b60168cc5e497",
	}

	dingdingToken = tokenMap[projectGroup]
	ossUrl := fmt.Sprintf("http://%s.oss-cn-hangzhou.aliyuncs.com/heapdump/%s/%s/%s-%s", bucketName, env, folder, podId, postfix) //建议OSS内网，"oss-cn-shanghai-internal"更换成自己的endpoint[需要修改]
	alarmMsg := fmt.Sprintf("<font color=#FF0000 size=5 face='黑体'>事故警告: JVM OOM</font>\n### 服务名: %s\n### podId: %s\n### 时间: %s\n### Dump文件: %s", projectName, podId, time.Now().Format("2006/01/02 15:04:05"), ossUrl)
	msgSender(dingdingToken, alarmMsg)
}
