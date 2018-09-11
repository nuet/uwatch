package service

import (
	"time"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"log"
	"strings"
)

type CmdbApi struct {
	requestType string
}

type CommonResponse struct {
	Result  bool        `json:"result"`
	Code    int         `json:"bk_error_code"`
	Message string      `json:"bk_error_msg"`
}

func (this *CmdbApi) SetRequestType(requestType string) {
	this.requestType = requestType
}

func (this *CmdbApi) Send(method string, row []byte) (string, error) {
	cmdbApiUrl := beego.AppConfig.String("CMDB_API")
	requestUrl := cmdbApiUrl + this.requestType
	var req *httplib.BeegoHTTPRequest
	var reqString string
	log.Println("api url:", requestUrl)
	switch strings.ToUpper(method) {
	case "GET":
		req = httplib.Get(requestUrl)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("BK_USER", "0")
		req.Header("HTTP_BLUEKING_SUPPLIER_ID", "0")
		reqstr, err := req.String()
		if err != nil {
			log.Println("cmdb api request fail", err)
			return reqstr, err
		} else {
			reqString = reqstr
		}
		break
	case "POST":
		req = httplib.Post(requestUrl)
		req.Body(row)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json")
		req.Header("BK_USER", "0")
		req.Header("HTTP_BLUEKING_SUPPLIER_ID", "0")
		reqstr, err := req.String()
		if err != nil {
			log.Println("cmdb api request fail", err)
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
		req.Header("BK_USER", "0")
		req.Header("HTTP_BLUEKING_SUPPLIER_ID", "0")
		reqstr, err := req.String()
		if err != nil {
			log.Println("cmdb api request fail", err)
			return reqstr, err
		} else {
			reqString = reqstr
		}
		break
	case "DELETE":
		req = httplib.Delete(requestUrl)
		req.SetTimeout(30*time.Second, 30*time.Second)
		req.Header("Content-Type", "application/json;charset=UTF-8")
		req.Header("BK_USER", "0")
		req.Header("HTTP_BLUEKING_SUPPLIER_ID", "0")

		reqstr, err := req.String()
		if err != nil {
			log.Println("cmdb api request fail", err)
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
