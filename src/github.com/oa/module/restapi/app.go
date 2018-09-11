package restapi

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sort"
)

const restHost = "http://restapi.juanpi.jp"

type App struct {
	Appkey string
	Secret string
}

type Values url.Values

type formValue []string

/**
 * 微信消息推送
 */
func NewWechatApp() *App {
	return &App{
		"oa",
		"!@oaapis&*#",
	}
}

/**
 * 短信消息推送
 */
func NewSmsApp() *App {
	return &App{
		"smsgate",
		"*_smsGate#-*(><]%^*AjT6@",
	}
}

/**
 * POST请求接口
 * @param  {[type]} app *App)         Post(path string, data url.Values [description]
 * @return {[type]}     [description]
 */
func (app *App) Post(path string, data url.Values) {
	data["appKey"] = formValue{app.Appkey}
	data["sign"] = formValue{generateSign(app.Secret, data)}
	resp, err := http.PostForm(restHost+path, data)
	if err != nil {
		io.WriteString(os.Stdout, err.Error())
	} else {
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

/**
 * POST请求提交JSON数据
 */
func (app *App) PostJSON(path string, data interface{}) {
	postData, err := json.Marshal(data)
	if err != nil {
		io.WriteString(os.Stdout, err.Error())
	}

	params := url.Values{}
	params["appKey"] = formValue{app.Appkey}
	requestURL := restHost + path + "?appKey=" + app.Appkey + "&sign=" + generateSign(app.Secret, params)
	resp, err := http.Post(requestURL, "application/json; charset=utf-8", bytes.NewReader(postData))
	if err != nil {
		io.WriteString(os.Stdout, err.Error())
	} else {
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

/**
 * 请求参数转换为字符串，参考url.Values.Encoding()
 */
func (v Values) toString() string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := k + "="
		for _, v := range vs {
			buf.WriteString(prefix)
			buf.WriteString(v)
		}
	}
	return buf.String()
}

/**
 * 根据请求参数生成签名
 */
func generateSign(secret string, v url.Values) string {
	varString := Values(v).toString()
	return fmt.Sprintf("%x", md5.Sum([]byte(varString+secret)))
}
