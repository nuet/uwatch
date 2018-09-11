package aliyunComm

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

type ErrorMessageResponse struct {
	XMLName   xml.Name `xml:"Error" json:"-"`
	Code      string   `xml:"Code,omitempty" json:"code,omitempty"`
	Message   string   `xml:"Message,omitempty" json:"message,omitempty"`
	RequestId string   `xml:"RequestId,omitempty" json:"request_id,omitempty"`
	HostId    string   `xml:"HostId,omitempty" json:"host_id,omitempty"`
}

type MessageResponse struct {
	XMLName   xml.Name `xml:"Message" json:"-"`
	Code      string   `xml:"Code,omitempty" json:"code,omitempty"`
	Message   string   `xml:"Message,omitempty" json:"message,omitempty"`
	RequestId string   `xml:"RequestId,omitempty" json:"request_id,omitempty"`
	HostId    string   `xml:"HostId,omitempty" json:"host_id,omitempty"`
}

type Base64Bytes []byte

type MessageReceiveResponse struct {
	MessageResponse
	MessageId        string      `xml:"MessageId" json:"message_id"`
	ReceiptHandle    string      `xml:"ReceiptHandle" json:"receipt_handle"`
	MessageBodyMD5   string      `xml:"MessageBodyMD5" json:"message_body_md5"`
	MessageBody      Base64Bytes `xml:"MessageBody" json:"message_body"`
	EnqueueTime      int64       `xml:"EnqueueTime" json:"enqueue_time"`
	NextVisibleTime  int64       `xml:"NextVisibleTime" json:"next_visible_time"`
	FirstDequeueTime int64       `xml:"FirstDequeueTime" json:"first_dequeue_time"`
	DequeueCount     int64       `xml:"DequeueCount" json:"dequeue_count"`
	Priority         int64       `xml:"Priority" json:"priority"`
}

type MessageBodyStruct struct {
	StatusCode string      `json:"status_code"`
	Duration   string      `json:"duration"`
	EndTime    string      `json:"end_time"`
	StatusMsg  string      `json:"status_msg"`
	StartTime  string      `json:"start_time"`
	OutID      string      `json:"out_id"`
	Dtmf       interface{} `json:"dtmf"`
	CallID     string      `json:"call_id"`
}

// RequestMNSQuery 为带有临时token认证客户端tmpService的方法
func (tmpService *Service) RequestMNSQuery() (response interface{}, err error) {

	tmpService.HTTPMethod = "GET"
	requestURL := MNSURL + MNSURI + "waitseconds=3"
	httpReq, err := http.NewRequest(tmpService.HTTPMethod, requestURL, nil)
	if err != nil {
		return nil, err
	}
	headerMap := tmpService.getHeader(MNSURI+"waitseconds=3", tmpService.Token)

	for k, v := range headerMap {
		httpReq.Header.Set(k, v)
	}

	httpResp, err := tmpService.HttpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	// println("statusCode: ", strconv.Itoa(httpResp.StatusCode))

	if httpResp.StatusCode > 204 {
		errResponse := ErrorMessageResponse{}
		err = xml.Unmarshal(body, &errResponse)
		if err != nil {
			return nil, err
		}
		return errResponse, nil

	}

	msgResponse := MessageReceiveResponse{}
	err = xml.Unmarshal(body, &msgResponse)
	if err != nil {
		return nil, err
	}

	return msgResponse, nil

}

// DeleteMessage 根据获取的ReceiptHandle 删除队列中对应消息
func (tmpService *Service) DeleteMessage(receiptHandle string) error {

	tmpService.HTTPMethod = "DELETE"
	beego.Debug("tmpService token: ", tmpService.Token)

	requestURL := MNSURL + MNSURI + "ReceiptHandle=" + receiptHandle

	httpReq, err := http.NewRequest(tmpService.HTTPMethod, requestURL, nil)
	headerMap := tmpService.getHeader(MNSURI+"ReceiptHandle="+receiptHandle, tmpService.Token)

	for k, v := range headerMap {
		httpReq.Header.Set(k, v)
	}

	httpResp, err := tmpService.HttpClient.Do(httpReq)
	if err != nil {
		return err
	}

	defer httpResp.Body.Close()

	beego.Debug("requestID of delete operation: ", httpResp.Header.Get("x-mns-request-id"))
	beego.Debug("requestID of delete operation: ", httpResp)

	return nil
}

// GetTmpService init a Service with tmpToken info, called before request MNS
func GetTmpService(trueService *Service) (tmpService *Service) {

	requestParams := TokenRequestParams{
		MessageType: "VoiceReport",
		RegionId:    "cn-hangzhou",
		Version:     "2017-05-25",
		Action:      "QueryTokenForMnsQueue",
	}

	var tokenResp TokenRespParams
	err := trueService.DoRequest(requestParams, &tokenResp)
	if err != nil {
		beego.Error("error : ", err)
	}

	keyID := tokenResp.MessageTokenDTO.AccessKeyId
	keySecrect := tokenResp.MessageTokenDTO.AccessKeySecret
	token := tokenResp.MessageTokenDTO.SecurityToken

	timeExpire, _ := time.Parse("2006-01-02 15:04:05", tokenResp.MessageTokenDTO.ExpireTime)

	keyPair := &AuthKeyPair{
		AccessKeyId:     keyID,
		AccessKeySecret: keySecrect,
	}

	aliyunD := DynamicConfig{
		Config:     keyPair,
		HTTPMethod: trueService.HTTPMethod,
		Timestamp:  trueService.Timestamp,
		Token:      token,
		Duration:   timeExpire,
	}

	tmpService = NewTokenService(&aliyunD)
	beego.Debug("KeyID: ", keyID, " KeySecrect: ", keySecrect, " Token: ", token)

	return

}

// IfExpire judge whether tmpService need to be refreshed
func (tmpService *Service) IfExpire(trueService *Service) (newService *Service) {

	beego.Debug("临时token 超时时间(据说是UTC时间): ", tmpService.Duration)

	timeNow := time.Now().Add(time.Hour * 8).UTC()

	beego.Debug("当前时间: ", timeNow)

	subMinute := tmpService.Duration.Sub(timeNow)

	beego.Debug("时间差 分钟数: ", subMinute.Minutes())

	beego.Debug(subMinute.Minutes())

	if int(subMinute.Minutes()) < 2 {
		newService = GetTmpService(trueService)
		beego.Info("token has been refreshed!")
		return
	}
	newService = tmpService
	return

}
