package captcha

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/env"
	"testing"
)

func init()  {
	env.SetupConfigEnv("../config.json")
}

func TestCreateTask(t *testing.T) {
	InstallCaptcha()
	key := GetCaptchaInstance().RandomCaptcha()
	fmt.Println(key)
	task, _ := GetCaptchaInstance().ImageToText(key.Image)
	if task !=nil {
		fmt.Println(task.Solution.Text)
	}
}