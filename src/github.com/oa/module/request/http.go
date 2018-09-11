package request

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"net/url"
)
var userAgent = ""
func SetUserAgent(ua string) {
	userAgent = ua
}

/**
 * POST请求接口
 */
func Post(path string, data map[string][]string, resData interface{}) (err error) {
	postData := url.Values(data)
	req, err := http.NewRequest("POST", path, strings.NewReader(postData.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去
	req.Header.Add("User-Agent", userAgent)
	client := &http.Client{}
	resp, err := client.Do(req)
	//resp, err := client.PostForm(path, data)
	if err != nil {
		log.Println(err)
		return
	} else {
		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(resData)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(resData)
		return
	}
}

func Get(url string, resData interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	} else {
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(resData)
		if err != nil {
			log.Println(err)
		}
		return
	}
}
