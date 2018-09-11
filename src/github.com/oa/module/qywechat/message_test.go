package qywechat

import "testing"

const notify = "huali"

var app = NewApp(34)

var testData = []string{
	"test & and",
	"有一个工作流【印章使用申请(NO:89210000020160805100724)彭帆】需要您进行审批，<a href='https://oa.juanpi.org/index.php?m=wap&c=workclass&a=view&workid=33365&type=2'>请点击进行审批</a>",
}

// 测试基本的消息发送
func TestSendText(t *testing.T) {
	for _, data := range testData {
		app.SendText(notify, data)
	}
}
