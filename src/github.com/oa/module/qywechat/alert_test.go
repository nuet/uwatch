package qywechat

import (
	"testing"
	"time"
)

func TestAlert(t *testing.T) {
	for i := 0; i <= 50; i++ {
		<-time.After(time.Second)
		Alert("huali", "test", "这是一个测试啊")
	}
}
