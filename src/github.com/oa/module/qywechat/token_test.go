package qywechat

import (
	"testing"

	"encoding/json"
)

var tests = map[int]string{
	34: "cy1EXgqEinbKvU3gr2y0_cvd15NiGq77pEnHjWRuSJh1Xujp_D2Z-S_3xSOU2qcY",
}

// func TestGetToken(t *testing.T) {
// 	for k, v := range tests {
// 		app := NewApp(k)
// 		if app.AccessToken.GetToken() != v {
// 			t.Error(v)
// 		}
// 	}
// }

// 测试解析json串是否正确
func TestAccessTokenParse(t *testing.T) {
	body := `{"access_token":"cy1EXgqEinbKvU3gr2y0_cvd15NiGq77pEnHjWRuSJh1Xujp_D2Z-S_3xSOU2qcY","expires_in":7200}`
	accessToken := &AccessToken{}
	json.Unmarshal([]byte(body), accessToken)
	if accessToken.Token != "cy1EXgqEinbKvU3gr2y0_cvd15NiGq77pEnHjWRuSJh1Xujp_D2Z-S_3xSOU2qcY" {
		t.Error(accessToken.Token)
	}
}
