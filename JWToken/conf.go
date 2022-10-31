package JWToken

import (
	"fmt"
	"github.com/spf13/viper"
)

type JWTConfig struct {
	SecretKey       string `json:"secret_key,omitempty"`
	AccessTokenTTL  int    `json:"access_token_ttl,omitempty"`
	SecureTokenTTL  int    `json:"secure_token_ttl,omitempty"`
	RefreshTokenTTL int    `json:"refresh_token_ttl,omitempty"`
}

var conf *JWTConfig

func getJWTConfigFromEnv() {
	conf = &JWTConfig{}

	if err := viper.UnmarshalKey("JWToken", conf); err != nil {
		err = fmt.Errorf("not found config name with env %q for JWToken with error: %+v", "JWToken", err)
		panic(err)
	}
}
