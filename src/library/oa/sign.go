package oa

import (
	"sort"
	"strings"
	"library/common"
)

// MakeSign 生成sign
func MakeSign(params map[string]string, secret string) string {
	keys := []string{}
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	str := ""
	for _, v := range keys {
		str += v + "=" + params[v]
	}
	str = common.Md5String(str + secret)
	sign := strings.ToLower(str)
	return sign
}