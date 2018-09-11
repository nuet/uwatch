package service

import (
	"time"
	"library/common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"log"
	"strings"
	"github.com/astaxie/beego/cache"
	"encoding/json"
	"fmt"
)


var GetCache cache.Cache

func init() {
	GetCache, _ = cache.NewCache("memory", `{"interval":60}`)
}

type JumpserverApi struct {
	requestType string
}

type LoginResponse struct {
	Token   string `json:"Token,omitempty"`
	Keyword string `json:"Keyword,omitempty"`
}

func (this *JumpserverApi) SetRequestType(requestType string) {
	this.requestType = requestType
}

func (this *JumpserverApi) Send(method string, row []byte) (string, error) {
	JumpserverApiUrl := beego.AppConfig.String("JUMPERSERVER_API")
	requestUrl := JumpserverApiUrl + this.requestType
	apiToken, err := this.RefreshToken()
	if err != nil {
		log.Println("Jumpserver token request fail", err)
		return apiToken, err
	}

	var req *httplib.BeegoHTTPRequest
	var reqString string
	log.Println("api url:", requestUrl)
	switch strings.ToUpper(method) {
	case "GET":
		req = httplib.Get(requestUrl)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("Authorization", apiToken)
		reqstr, err := req.String()
		if err != nil {
			log.Println("Jumpserver api request fail", err)
			return reqstr, err
		}
		resp, err := req.Response()
		if err != nil {
			log.Println("Jumpserver api response ", reqstr, " is err: ", err)
			return reqstr, err
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			beego.Error("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
			return reqstr, fmt.Errorf("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
		}
		reqString = reqstr
		break
	case "POST":
		req = httplib.Post(requestUrl)
		req.Body(row)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json")
		req.Header("Authorization", apiToken)
		reqstr, err := req.String()
		if err != nil {
			log.Println("Jumpserver api request fail", err)
			return reqstr, err
		}
		resp, err := req.Response()
		if err != nil {
			log.Println("Jumpserver api response ", reqstr, " is err: ", err)
			return reqstr, err
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			beego.Error("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
			return reqstr, fmt.Errorf("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
		}
		reqString = reqstr
		break
	case "PUT":
		req = httplib.Put(requestUrl)
		req.Body(row)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("Authorization", apiToken)
		reqstr, err := req.String()
		if err != nil {
			log.Println("Jumpserver api request fail", err)
			return reqstr, err
		}
		resp, err := req.Response()
		if err != nil {
			log.Println("Jumpserver api response ", reqstr, " is err: ", err)
			return reqstr, err
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			beego.Error("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
			return reqstr, fmt.Errorf("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
		}
		reqString = reqstr
		break
	case "DELETE":
		req = httplib.Delete(requestUrl)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("Authorization", apiToken)
		reqstr, err := req.String()
		if err != nil {
			log.Println("Jumpserver api request fail", err)
			return reqstr, err
		}
		resp, err := req.Response()
		if err != nil {
			log.Println("Jumpserver api response ", reqstr, " is err: ", err)
			return reqstr, err
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			beego.Error("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
			return reqstr, fmt.Errorf("Jumpserver api error code ", resp.StatusCode, " resp: ", reqstr)
		}
		reqString = reqstr
		break
	default:
		log.Println("method error", method)
	}
	if reqString == "" {
		log.Println("http rsp is null")
	}
	return reqString, nil
}

func (this *JumpserverApi) RefreshToken() (string, error) {
	var apiToken string
	if !GetCache.IsExist("JumpserverApiToken") || common.GetString(GetCache.Get("JumpserverApiToken")) == "" {
		loginUrl := beego.AppConfig.String("JUMPERSERVER_API") + "/api/users/v1/token/"
		data := map[string]string{
			"username" : "admin",
			"password" : "admin",
		}
		row, _ := json.Marshal(data)

		var req *httplib.BeegoHTTPRequest
		log.Println("api url:", loginUrl)
		req = httplib.Post(loginUrl)
		req.Body(row)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json")
		reqstr, err := req.String()
		if err != nil {
			log.Println("Jumpserver api request fail", err)
			return apiToken, err
		}
		resp, err := req.Response()
		if err != nil {
			log.Println("Jumpserver token response ", reqstr, " is err: ", err)
			return apiToken, err
		}
		if resp.StatusCode < 200 || resp.StatusCode >= 400 {
			beego.Error("Jumpserver token error code ", resp.StatusCode, " resp: ", reqstr)
			return apiToken, fmt.Errorf("Jumpserver token error code ", resp.StatusCode, " resp: ", reqstr)
		}

		rstObj := &LoginResponse{}
		jsErr := json.Unmarshal([]byte(reqstr), rstObj);
		if nil != jsErr {
			fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
			return apiToken, jsErr
		}

		apiToken = rstObj.Keyword + " " + rstObj.Token
		GetCache.Put("JumpserverApiToken", apiToken, 3000 * time.Second)
	} else {
		apiToken = common.GetString(GetCache.Get("JumpserverApiToken"))
	}

	return apiToken, nil
}