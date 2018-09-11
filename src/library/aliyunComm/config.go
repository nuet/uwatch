package aliyunComm

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	// "strings"
)

const (
	BASEURL          = "http://dyvmsapi.aliyuncs.com/?"
	MNSQUERYNAME     = "Alicom-Queue-1650578505207239-VoiceReport"
	SIGNATUREVERSION = "1.0"
	TOKENURL         = "http://dybaseapi.aliyuncs.com/?"
	MNSURL           = "https://1943695596114318.mns.cn-hangzhou.aliyuncs.com"
	MNSURI           = "/queues/Alicom-Queue-1650578505207239-VoiceReport/messages?"
)

type AuthKeyPair struct {
	AccessKeyId     string
	AccessKeySecret string
}

type CommonRequest struct {
	AccessKeyId      string
	AccessKeySecret  string
	Timestamp        string
	Format           string
	SignatureMethod  string
	SignatureVersion string
	SignatureNonce   string
	Signature        string
}

type DynamicConfig struct {
	Config          *AuthKeyPair
	HTTPMethod      string
	Timestamp       string
	Format          string
	SignatureMethod string
	SignatureNonce  string
	Token           string
	Duration        time.Time
}

type Service struct {
	Config           *AuthKeyPair
	HTTPMethod       string
	Timestamp        string
	Format           string
	SignatureMethod  string
	SignatureVersion string
	SignatureNonce   string
	Signature        string
	Token            string
	Duration         time.Time

	BaseUrl    string
	HttpClient *http.Client
}

func (s *Service) SetURL(params interface{}) (string, error) {
	if len(s.BaseUrl) == 0 {
		return "", errors.New("BaseUrl is not set")
	}

	commonRequest := CommonRequest{
		AccessKeyId:      s.Config.AccessKeyId,
		Timestamp:        time.Now().UTC().Format(time.RFC3339),
		Format:           s.Format,
		SignatureMethod:  s.SignatureMethod,
		SignatureVersion: s.SignatureVersion,
		SignatureNonce:   uuid.NewV4().String(),
	}

	values := url.Values{}

	ConvertParamsToValues(commonRequest, &values)
	ConvertParamsToValues(params, &values)

	a, err := GetURL(values, s.HTTPMethod, s.Config.AccessKeySecret)
	if err != nil {
		return "", err
	}
	reqURL := s.BaseUrl + a
	return reqURL, nil
}

func ConvertParamsToValues(params interface{}, values *url.Values) {
	elem := reflect.ValueOf(params)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}

	elemType := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		fieldName := elemType.Field(i).Name

		field := elem.Field(i)
		kind := field.Kind()
		if (kind == reflect.Ptr ||
			kind == reflect.Array ||
			kind == reflect.Slice ||
			kind == reflect.Map ||
			kind == reflect.Chan) && field.IsNil() {
			continue

		}

		if kind == reflect.Ptr {
			field = field.Elem()
			kind = field.Kind()
		}

		var v string
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if field.Int() != 0 {
				v = strconv.FormatInt(field.Int(), 10)
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if field.Uint() != 0 {
				v = strconv.FormatUint(field.Uint(), 10)
			}

		case reflect.Float32:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 32)

		case reflect.Float64:
			v = strconv.FormatFloat(field.Float(), 'f', 4, 64)

		case reflect.Bool:
			v = strconv.FormatBool(field.Bool())

		case reflect.String:
			v = field.String()
		case reflect.Slice:
			switch field.Type().Elem().Kind() {
			case reflect.String:
				l := field.Len()
				if l > 0 {
					for i := 0; i < l; i++ {
						v = field.Index(i).String()
						if v != "" {
							//name := elemType.Field(i).Tag.Get("ArgName")
							//if name == "" {
							name := fieldName
							//}
							name = fmt.Sprintf("%s.%d", name, i)
							values.Set(name, v)
						}
					}
					continue
				}
			default:

			}
		}

		if v != "" {
			name := elemType.Field(i).Tag.Get("ArgName")
			if name == "" {
				name = fieldName
			}
			values.Set(name, v)
		}
	}
}

func GetURL(values url.Values, method, keySecret string) (string, error) {

	if values == nil {
		return "", fmt.Errorf("values is empty")
	}
	var strToSig, urlEncode string
	params := values.Encode()
	urlEncode = url.QueryEscape(params)
	// fmt.Println("values encode : ", urlEncode)

	strToSig = method + "&" + url.QueryEscape("/") + "&" + urlEncode

	// beego.Debug("待签名字串: ", strToSig)
	// beego.Debug("secret key :", keySecret)

	sigEncode := sign(strToSig, keySecret)
	urlToRequest := "Signature=" + sigEncode + "&" + params

	return urlToRequest, nil
}

func sign(strToSign string, secretKey string) string {

	key := []byte(secretKey + "&")

	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(strToSign))
	sig := base64.StdEncoding.EncodeToString([]byte(string(hash.Sum(nil))))
	encdSig := url.QueryEscape(sig)
	return encdSig

}
