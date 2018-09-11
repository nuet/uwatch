## 卷皮内部基础模块调用

	require  golang 1.7

### 调用企业微信服务
	cd qywechat

	// 测试发送企业微信消息
	go test --run SendText

### 调用restapi服务

	cd restapi
	// 测试微信企业消息推送
	go test --run PushWechat

	// 测试短信消息推送
	go test --run PushSMS  

	// 测试消息中心消息推送
	go test --run MsgCenter