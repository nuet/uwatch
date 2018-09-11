package qywechat

type Login struct {
	OpenId   int
	UserId   string
	DeviceId int
}

const (
	METHOD_GETINFO = "/user/getuserinfo"
)

// GetInfo 获取用户信息
func (l *Login) GetInfo(code string) error {
	// 默认为oa应用id
	app := NewApp(0)
	data := map[string]string{"code": code}
	return app.Get(METHOD_GETINFO, data, l)
}
