package qywechat

import (
	"fmt"
	"testing"
)

// 测试基本的消息发送
func TestGetInfo(t *testing.T) {
	l := Login{}
	err := l.GetInfo("aaa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(l)
}
