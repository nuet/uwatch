package qywechat

import (
	"fmt"
	"github.com/oa/module/request"
	"reflect"
)

type App struct {
	CorpId      string
	CorpSecret  string
	AgentId     int64
	AccessToken *AccessToken
}

const (
	baseUrl    = "https://qyapi.weixin.qq.com/cgi-bin"
	corpId     = "wx79563e5ff4b91692"
	corpSecret = "Hd30w9gts4eS53v5sAgfZYaVnTgluV936gOnm_ZVnA0TYCxzwUIVyipMOmXBy-MT"
)

// NewApp 创建一个新的App agentId就是企业号的应用ID
func NewApp(agentId int64) *App {
	return &App{
		corpId,
		corpSecret,
		agentId,
		&AccessToken{},
	}
}

// Post App发起Post请求
func (app *App) Post(path string, data interface{}, response interface{}) (err error) {
	requestURL := baseUrl + path + "?access_token="
	hasRetried := false
	var url string
RETRY:
	url = requestURL + app.GetToken()
	err = request.PostJSON(url, data, &response)
	if err != nil {
		return
	}

	var ErrorStructValue reflect.Value // Error
	responseStructValue := reflect.ValueOf(response).Elem()
	if v := responseStructValue.Field(0); v.Kind() == reflect.Struct {
		ErrorStructValue = v
	} else {
		ErrorStructValue = responseStructValue
	}

	switch ErrCode := ErrorStructValue.Field(0).Int(); ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeAccessTokenExpired:
		if !hasRetried {
			hasRetried = true
			err = app.RequestToken()
			if err != nil {
				return
			}
			responseStructValue.Set(reflect.New(responseStructValue.Type()).Elem())
			goto RETRY
		}
		fallthrough
	default:
		// 如果发送失败，则把发送的内容也记录下来
		fmt.Printf("errcode:%s, %v\n", ErrCode, data)
		return
	}

	return
}

// Get App发起Get请求
func (app *App) Get(path string, query map[string]string, response interface{}) (err error) {
	requestURL := baseUrl + path + "?access_token="
	queryStr := ""
	for k, v := range query {
		queryStr += "&" + k + "=" + v
	}

	hasRetried := false
	var url string
RETRY:
	url = requestURL + app.GetToken() + queryStr
	err = request.GetJSON(url, &response)
	if err != nil {
		return
	}
	var ErrorStructValue reflect.Value // Error
	responseStructValue := reflect.ValueOf(response).Elem()
	if v := responseStructValue.Field(0); v.Kind() == reflect.Struct {
		ErrorStructValue = v
	} else {
		ErrorStructValue = responseStructValue
	}

	switch ErrCode := ErrorStructValue.Field(0).Int(); ErrCode {
	case ErrCodeOK:
		return
	case ErrCodeAccessTokenExpired:
		if !hasRetried {
			hasRetried = true
			err = app.RequestToken()
			if err != nil {
				return
			}
			responseStructValue.Set(reflect.New(responseStructValue.Type()).Elem())
			goto RETRY
		}
		fallthrough
	default:
		return
	}

	return
}
