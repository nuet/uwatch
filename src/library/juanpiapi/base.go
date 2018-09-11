package juanpiapi

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"github.com/oa/module/request"
)

type Api struct {
	Url    string
	Appkey string
	Secret string
}

type RespData struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
}

func (e *RespData) Error() string {
	return e.Msg
}

func (api *Api) Post(path string, data map[string]string, result interface{}) (err error) {

	postdata := map[string][]string{}
	for k, v := range data {
		postdata[k] = []string{v}
	}

	url := api.getRequestUrl(path, data)
	request.SetUserAgent("chaoge")
	err = request.Post(url, postdata, result)
	if err != nil {
		fmt.Println(url)
		fmt.Println("请求失败，" + err.Error())
	}
	return
}

func (api *Api) getRequestUrl(path string, data map[string]string) string {
	return api.Url + api.getKeySignQuery(path, data)
}

func (api *Api) getKeySignQuery(path string, data map[string]string) string {
	sign := api.getApiSign(data)
	return fmt.Sprintf("%s?appKey=%s&sign=%s", path, api.Appkey, sign)
}

func (api *Api) getApiSign(data map[string]string) string {
	sign := "appKey=" + api.Appkey
	sorted_keys := make([]string, 0)
	if len(data) > 0 {
		for k, _ := range data {
			sorted_keys = append(sorted_keys, k)
		}
		sort.Strings(sorted_keys)
		for _, k := range sorted_keys {
			sign += fmt.Sprintf("%v=%v", k, data[k])
		}
	}
	sign += api.Secret
	c := md5.New()
	c.Write([]byte(sign)) // 需要加密的字符串为 sign
	cipherStr := c.Sum(nil)
	return hex.EncodeToString(cipherStr) // 加密结果
}
