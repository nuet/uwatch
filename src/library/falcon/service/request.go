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
	"github.com/pkg/errors"
)

var GetCache cache.Cache

func init() {
	GetCache, _ = cache.NewCache("memory", `{"interval":60}`)
}

type FalconApi struct {
	requestType string
	requestLogin string
}

type CommonResponse struct {
	ErrorMsg  string     	`json:"error"`
}

type CommonSuccResponse struct {
	SuccMsg   string  	`json:"message"`
}

type LoginResponse struct {
	CommonResponse
	Sig   string	    	`json:"sig"`
	Name  string		`json:"name"`
	Admin bool		`json:"admin"`
}

func (this *FalconApi) SetRequestType(requestType string) {
	this.requestType = requestType
}

func (this *FalconApi) SetRequestLogin(requestLogin string) {
	this.requestLogin = requestLogin
}

func (this *FalconApi) Send(method string, row []byte) (string, error) {
	FalconApiUrl := beego.AppConfig.String("FALCON_API")
	requestUrl := FalconApiUrl + this.requestType
	requestLogin :=  this.requestLogin
	apiToken := this.RefreshToken(requestLogin)
	var req *httplib.BeegoHTTPRequest
	var reqString string
	log.Println("api url:", requestUrl)
	switch strings.ToUpper(method) {
	case "GET":
		req = httplib.Get(requestUrl)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("Apitoken", apiToken)
		s, _ := req.Response()
		if s.StatusCode != 200 {
			rsp := &CommonResponse{}
			err := req.ToJSON(rsp)
			if err != nil {
				log.Println("Falcon api request fail", err)
				return "", err
			} else if rsp.ErrorMsg != "" {
				log.Println("Falcon api request fail", rsp.ErrorMsg)
				return "", errors.New(rsp.ErrorMsg)
			}
		}
		reqstr, err := req.String()
		if err != nil {
			log.Println("Falcon api request fail", err)
			return reqstr, err
		} else {
			reqString = reqstr
		}
		break
	case "POST":
		req = httplib.Post(requestUrl)
		req.Body(row)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("Apitoken", apiToken)
		s, _ := req.Response()
		if s.StatusCode != 200 {
			rsp := &CommonResponse{}
			err := req.ToJSON(rsp)
			if err != nil {
				log.Println("Falcon api request fail", err)
				return "", err
			} else if rsp.ErrorMsg != "" {
				log.Println("Falcon api request fail", rsp.ErrorMsg)
				return "", errors.New(rsp.ErrorMsg)
			}
		}
		reqstr, err := req.String()
		if err != nil {
			log.Println("Falcon api request fail", err)
			return reqstr, err
		} else {
			reqString = reqstr
		}
		break
	case "PUT":
		req = httplib.Put(requestUrl)
		req.Body(row)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("Apitoken", apiToken)
		s, _ := req.Response()
		if s.StatusCode != 200 {
			rsp := &CommonResponse{}
			err := req.ToJSON(rsp)
			if err != nil {
				log.Println("Falcon api request fail", err)
				return "", err
			} else if rsp.ErrorMsg != "" {
				log.Println("Falcon api request fail", rsp.ErrorMsg)
				return "", errors.New(rsp.ErrorMsg)
			}
		}
		reqstr, err := req.String()
		if err != nil {
			log.Println("Falcon api request fail", err)
			return reqstr, err
		} else {
			reqString = reqstr
		}
		break
	case "DELETE":
		req = httplib.Delete(requestUrl)
		req.Body(row)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("Apitoken", apiToken)
		s, _ := req.Response()
		if s.StatusCode != 200 {
			rsp := &CommonResponse{}
			err := req.ToJSON(rsp)
			if err != nil {
				log.Println("Falcon api request fail", err)
				return "", err
			} else if rsp.ErrorMsg != "" {
				log.Println("Falcon api request fail", rsp.ErrorMsg)
				return "", errors.New(rsp.ErrorMsg)
			}
		}
		reqstr, err := req.String()
		if err != nil {
			log.Println("Falcon api request fail", err)
			return reqstr, err
		} else {
			reqString = reqstr
		}
		break
	default:
		log.Println("method error", method)
	}
	if reqString == "" {
		log.Println("http rsp is null")
	}
	return reqString, nil
}

func (this *FalconApi) RefreshToken(login string) string {
	var apiToken string
	if login == "" {
		if !GetCache.IsExist("FalconApiToken") || common.GetString(GetCache.Get("FalconApiToken")) == "" {
			loginUrl := beego.AppConfig.String("FALCON_API") + "/user/login"
			data := map[string]string{
				"name" : "root",
				"password" : "itil@juanpi.com",
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
				log.Println("Falcon api request fail", err)
				return ""
			}
			rstObj := &LoginResponse{}
			jsErr := json.Unmarshal([]byte(reqstr), rstObj);
			if nil != jsErr {
				fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
				return ""
			}
			if rstObj.ErrorMsg != "" {
				fmt.Println("failed to unmarshal the result, error info is ", rstObj.ErrorMsg)
				return ""
			}
			rsp := map[string]string{
				"name" : rstObj.Name,
				"sig"  : rstObj.Sig,
			}
			bi, _ := json.Marshal(rsp)
			apiToken = string(bi)
			GetCache.Put("FalconApiToken", apiToken, 2419200 * time.Second)
		} else {
			apiToken = common.GetString(GetCache.Get("FalconApiToken"))
		}
	} else {
			loginUrl := beego.AppConfig.String("FALCON_API") + "/user/login"
			data := map[string]string{
				"name" : login,
				"password" : login,
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
				log.Println("Falcon api request fail", err)
				return ""
			}
			rstObj := &LoginResponse{}
			jsErr := json.Unmarshal([]byte(reqstr), rstObj);
			if nil != jsErr {
				fmt.Println("failed to unmarshal the result, error info is ", jsErr.Error())
				return ""
			}
			if rstObj.ErrorMsg != "" {
				fmt.Println("failed to unmarshal the result, error info is ", rstObj.ErrorMsg)
				return ""
			}
			rsp := map[string]string{
				"name" : rstObj.Name,
				"sig"  : rstObj.Sig,
			}
			bi, _ := json.Marshal(rsp)
			apiToken = string(bi)
	}

	return apiToken
}