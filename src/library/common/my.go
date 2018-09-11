package common

import "time"

//获取相差分钟，负数为0
func GetMinDiffer(start_time, end_time string) int {
	if start_time == "" || end_time == "" {
		return 0
	}
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", start_time, time.Local)
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", end_time, time.Local)
	if err == nil && t1.Before(t2) {
		diff := t2.Unix() - t1.Unix()
		return int(diff) / 60
	} else {
		return 0
	}
}

//获取公司
func GetCompany(company string) string {
	zh_company := ""
	if company == "wh" {
		zh_company = "卷皮武汉厂"
	}
	if company == "sz" {
		zh_company = "卷皮深圳厂"
	}
	if company == "hz" {
		zh_company = "卷皮杭州厂"
	}
	if company == "bj" {
		zh_company = "卷皮北京厂"
	}
	return zh_company
}
