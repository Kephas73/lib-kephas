package JWToken

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/Kephas73/lib-kephas/env"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

func init() {
	env.SetupConfigEnv("../config.json")
}

var token1 = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfdXVpZCI6IjFkMmJmMzAwLTA5NWYtNDMwNi04MDllLTQ1ZDM0ZjFiOGVhNiIsImF1dGhvcml6ZWQiOnRydWUsImRldmljZV9pcCI6IjEyNy4wLjAuMSIsImRldmljZV9raW5kIjozLCJleHAiOjE2NjcyMzY2ODEsInNlY3VyZV9leHAiOjE2NjcyMzQ4ODEsInVzZXJfaWQiOiJkMTc0MDg4ZS0wM2YxLTQ0Y2MtYmYxZi02NWZhOTBiOTE1MmMifQ.YeC8N9Xn47RUd_3I2btm-ovAfColxMeEqUlQCsN0kxs`

func TestExtractToken(t *testing.T) {
	token, err := NewToken(base.NewUuid(), constant.KDeviceKindWeb, "127.0.0.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token.AccessToken)
	fmt.Println(token)

	token2, _ := jwt.Parse(token1, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Signing Key
		signKey := conf.SecretKey
		return []byte(signKey), nil
	})

	fmt.Println(token2)
}
