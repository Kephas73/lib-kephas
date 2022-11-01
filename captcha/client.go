package captcha

import (
	"fmt"
)

type CaptchaClient struct {
	config  *Config
	captcha []*Captcha
}

var captchaInstance *CaptchaClient

func InstallCaptcha(configKeys ...string) {

	if configCaptcha == nil {
		getConfigFromEnv(configKeys...)
	}

	if configCaptcha == nil {
		err := fmt.Errorf("need config for captcha client first")
		panic(err)
	}

	captchaInstance = &CaptchaClient{configCaptcha, NewCaptcha(configCaptcha.PathKey)}

	return
}

func GetCaptchaInstance() *CaptchaClient {
	if captchaInstance == nil {
		InstallCaptcha()
	}

	return captchaInstance
}
