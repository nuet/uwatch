package notice
import (
	"fmt"
	"library/aliyunComm"
	"models"
	"strings"
	"time"
	"library/oa"
	"errors"
	"github.com/astaxie/beego"
)


//定义语音告警对外接口的统一响应体
type VMsgApiReq struct {
	Mobile string `json:"mobile"`
	VMsg   string `json:"vMsg"`
	Type   string `json:"type"`
}

//定义语音告警对外接口的统一响应体
type VMsgApiResp struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

//通用语音报警通知
func SendAlertReport(message, mobile, templateid, actionid, alert_type string, count int) (voiceRecord models.UwVoicerecord, err error) {
	// getTmplateID 根据类型获取对应语音模板ID
	templateid = getTmplateID(alert_type)
	voiceRecord, err = SendAlertReportByAli(message, mobile, templateid, actionid, alert_type, count)
	return voiceRecord, nil
}

// SendAlertReportByAli 阿里云语音告警接口
func SendAlertReportByAli(message, mobile, templateid, actionid, alert_type string, count int) (voiceRecord models.UwVoicerecord, err error) {

	if mobile == "" {
		beego.Error("号码为空")
		return
	}
	message = strings.Replace(message, " ", "", -1)

	if actionid == "" {
		actionid = "m" + mobile + "t" + time.Now().Format("20060102150405")
	}
	userInfo, err := oa.GetUser(mobile)
	if err != nil {
		beego.Info("Getuser error!! ", err)
	}
	user := "未知用户"
	if len(userInfo.Users) > 0 {
		user = userInfo.Users[0].CnName + "." + userInfo.Users[0].Name
	}

	keyPair := &aliyunComm.AuthKeyPair{
		AccessKeyId:     "LTAIKdNSARLkQfC6",
		AccessKeySecret: "Fby1anRqSQ3mF5imaRyvgONkgIFYcP",
	}

	aliyunD := aliyunComm.DynamicConfig{
		Config:          keyPair,
		HTTPMethod:      "GET",
		Timestamp:       time.Now().UTC().Format(time.RFC3339),
		Format:          "JSON",
		SignatureMethod: "HMAC-SHA1",
	}
	aliyunService := aliyunComm.New(&aliyunD)

	var resp aliyunComm.Response
	requestParams := aliyunComm.RequsetParams{
		Action:           "SingleCallByTts",
		Version:          "2017-05-25",
		CalledNumber:     mobile,
		CalledShowNumber: "02759769899",
		OutId:            actionid,
		RegionId:         "cn-hangzhou",
		TtsCode:          templateid,
		TtsParam:         fmt.Sprintf("{\"item\":\"%s\"}", message),
	}

	err = aliyunService.DoRequest(requestParams, &resp)
	if err != nil {
		beego.Error("aliyun 语音接口调用出错,", err.Error())
	}

	if resp.Code != "OK" {
		err = errors.New(resp.Message)
		beego.Error(err)
	}

	voiceRecord = models.UwVoicerecord{
		ActionId:      actionid,
		CallId:        resp.CallId,
		VoiceMessage:  message,
		Mobile:        mobile,
		AlertUser:     user,
		AlertType:     alert_type,
		Result:        resp.Code,
		FailureReason: resp.Message,
		Count:         count,
	}
	_, err = models.AddVoiceRecord(&voiceRecord)
	if err != nil {
		beego.Info("数据库写入错误:  ", err)
		return voiceRecord, err
	}
	//语音告警结果会回调
	return voiceRecord, nil
}

func getTmplateID(alertType string) (templateid string) {
	//获取阿里云语音模板ID
	switch alertType {
	case "OPS":
		templateid = "TTS_89995018"
	case "BI":
		templateid = "TTS_113175065"
	case "RD":
		templateid = "TTS_113180083"
	}
	return

}


