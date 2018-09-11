package request

import (
	"bytes"
	"encoding/json"
	"net/http"
)

/**
 * 请求接口并返回json格式的数据
 * @param {[type]} url string) (encoder *json.Encoder, err error [description]
 */
func GetJSON(url string, v interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(v)
		return err
	}
}

// PostJSON POST请求，同时将返回的json结果解析到result
func PostJSON(url string, data, result interface{}) error {
	var buf bytes.Buffer
	postEncoder := json.NewEncoder(&buf)

	postEncoder.SetEscapeHTML(false)
	err := postEncoder.Encode(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json; charset=utf-8", bytes.NewReader(buf.Bytes()))
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(result)
		return err
	}
}
