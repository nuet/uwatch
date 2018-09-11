package monitor

import "testing"

var testURL = []string{
	"http://www.baidu.com",
	"http://192.168.16.5",
	"http://192.168.16.9",
}

// 测试基本的消息发送
func TestCheckURL(t *testing.T) {
	for _, URL := range testURL {
		err := CheckURL(URL, 2)
		if err != nil {
			println(err.Error())
		}
	}
}
