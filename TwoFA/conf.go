package TwoFA

import (
	"fmt"
	"github.com/pquerna/otp"
	"github.com/spf13/viper"
	"strings"
)

type Env struct {
	Issuer     string `json:"issuer"`
	Period     uint   `json:"period"`
	SecretSize uint   `json:"secret_size"`
	Digits     int    `json:"digits"`
	Algorithm  string `json:"algorithm"`
}

type Config struct {
	Issuer     string
	Period     uint
	SecretSize uint
	Digits     otp.Digits
	Algorithm  otp.Algorithm
}

var env *Env
var twoFAConfig *Config

func getTwoFAFromEnv(configKeys ...string) {
	configKey := "2FA"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	env = &Env{}
	twoFAConfig = &Config{}
	if err := viper.UnmarshalKey(configKey, &env); err != nil {
		err = fmt.Errorf("not found config name with env %q for 2FA with error: %+v", configKey, err)
		panic(err)
	}

	if env.Issuer == "" {
		err := fmt.Errorf("not found isuer property of config for %q", configKey)
		panic(err)
	} else {
		twoFAConfig.Issuer = env.Issuer
	}

	if env.Period <= 0 {
		twoFAConfig.Period = 30
	} else {
		twoFAConfig.Period = env.Period
	}

	if env.SecretSize <= 0 {
		twoFAConfig.SecretSize = 20
	} else {
		twoFAConfig.SecretSize = env.SecretSize
	}

	switch env.Digits {
	case 6:
		twoFAConfig.Digits = otp.DigitsSix
	case 8:
		twoFAConfig.Digits = otp.DigitsEight
	default:
		twoFAConfig.Digits = otp.DigitsSix
	}

	switch env.Algorithm {
	case "SHA1":
		twoFAConfig.Algorithm = otp.AlgorithmSHA1
	case "SHA256":
		twoFAConfig.Algorithm = otp.AlgorithmSHA256
	case "SHA512":
		twoFAConfig.Algorithm = otp.AlgorithmSHA512
	case "MD5":
		twoFAConfig.Algorithm = otp.AlgorithmMD5
	default:
		twoFAConfig.Algorithm = otp.AlgorithmSHA1
	}
}
