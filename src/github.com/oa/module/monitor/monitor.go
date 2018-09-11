package monitor

import (
	"fmt"
	"net/http"
	"time"
)

/**
 * 监控url地址
 */
func CheckURL(url string, timeout time.Duration) error {
	// 等待8秒，否则超时
	timeoutSecond := time.Duration(timeout * time.Second)
	client := http.Client{
		Timeout: timeoutSecond,
	}
	fmt.Printf("%s request url: %s\n", time.Now().Format("2006-01-02 15:04:05"), url)
	resp, err := client.Get(url)
	if err != nil {
		return err
	} else {
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("%s访问返回状态码:%d", url, resp.StatusCode)
		}
	}

	return nil
}
