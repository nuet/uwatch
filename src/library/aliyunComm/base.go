package aliyunComm

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type TokenRequestParams struct {
	MessageType string
	RegionId    string
	Version     string
	Action      string
}

type TokenRespParams struct {
	MessageTokenDTO struct {
		AccessKeySecret string
		AccessKeyId     string
		ExpireTime      string
		CreateTime      string
		SecurityToken   string
	}
	Message   string
	RequestId string
	Code      string
}

func (tmpService *Service) getHeader(uri, tmpToken string) map[string]string {
	// headerMap["security-token"] = tmpService.Token

	headerMap := make(map[string]string)

	headerMap["Connection"] = "keep-alive"
	headerMap["x-mns-version"] = "2015-06-06"
	headerMap["Content-Type"] = "text/xml;charset=UTF-8"
	headerMap["Date"] = time.Now().UTC().Format(http.TimeFormat)
	headerMap["Host"] = MNSURL

	var keyList []string
	var canonicalizedMNSHeader, canonicalizedResource, stringToSig string

	for k := range headerMap {
		keyList = append(keyList, strings.ToLower(k))
	}
	sort.Strings(keyList)

	for _, key := range keyList {
		if strings.HasPrefix(key, "x-mns-") {
			canonicalizedMNSHeader = key + ":" + headerMap[key] + "\n"
		}
	}
	canonicalizedResource = uri

	stringToSig = fmt.Sprintf("%s\n%s\n%s\n%s\n%s%s", tmpService.HTTPMethod, "", headerMap["Content-Type"], headerMap["Date"], canonicalizedMNSHeader, canonicalizedResource)

	hash := hmac.New(sha1.New, []byte(tmpService.Config.AccessKeySecret))
	hash.Write([]byte(stringToSig))
	sig := base64.StdEncoding.EncodeToString([]byte(string(hash.Sum(nil))))
	beego.Debug("signature: ", sig)

	headerMap["Authorization"] = "MNS " + tmpService.Config.AccessKeyId + ":" + sig
	headerMap["security-token"] = tmpToken

	return headerMap
}

// DoRequest 账户Service方法, 用于请求语音接口
func (tmpService *Service) DoRequest(params interface{}, response interface{}) error {
	requestURL, err := tmpService.SetURL(params)
	beego.Debug(requestURL)
	if err != nil {
		return err
	}
	httpReq, err := http.NewRequest(tmpService.HTTPMethod, requestURL, nil)
	if err != nil {
		return err
	}

	httpResp, err := tmpService.HttpClient.Do(httpReq)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &response); err != nil {
		return err
	}

	return nil

}

func New(d *DynamicConfig) *Service {
	service := &Service{
		Config:           d.Config,
		Timestamp:        d.Timestamp,
		Format:           d.Format,
		SignatureMethod:  d.SignatureMethod,
		SignatureVersion: SIGNATUREVERSION,
		SignatureNonce:   d.SignatureNonce,
		HTTPMethod:       d.HTTPMethod,

		BaseUrl:    BASEURL,
		HttpClient: &http.Client{},
	}
	return service
}

func NewTokenService(d *DynamicConfig) *Service {

	service := &Service{
		Config:           d.Config,
		Timestamp:        d.Timestamp,
		Format:           d.Format,
		SignatureMethod:  d.SignatureMethod,
		SignatureVersion: SIGNATUREVERSION,
		HTTPMethod:       d.HTTPMethod,
		Token:            d.Token,
		Duration:         d.Duration,

		BaseUrl:    TOKENURL,
		HttpClient: &http.Client{},
	}
	return service
}
