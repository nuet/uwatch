package aliyunComm

import "github.com/astaxie/beego"

type RequsetParams struct {
	Action           string
	Version          string
	RegionId         string
	CalledShowNumber string
	CalledNumber     string
	TtsCode          string
	TtsParam         string
	OutId            string
}

type Response struct {
	RequestId string
	Code      string
	Message   string
	CallId    string
}

// SingleCallByTts :Using trueService to Request SingleCallByTts
func (trueService *Service) SingleCallByTts(requestP *RequsetParams) (resp *Response, err error) {
	err = trueService.DoRequest(requestP, &resp)
	return resp, nil
}

// GetMNSQuery :Using tmpService to request MNSquery after judging whether tmpService need to be refresh
func GetMNSQuery(tmpService, trueService *Service, respChan chan MessageReceiveResponse, errChan chan ErrorMessageResponse) {
	tmpService = tmpService.IfExpire(trueService)
	var resp interface{}

	resp, err := tmpService.RequestMNSQuery()
	if err != nil {
		beego.Debug(err)
		beego.Debug("error in GetMNSQuery!")
	}
	switch t := resp.(type) {
	case MessageReceiveResponse:
		// beego.Debug("getting MsgResp:")
		respChan <- t
		return
	case ErrorMessageResponse:
		errChan <- t
		return
	}

}
