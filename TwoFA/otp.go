package TwoFA

import (
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"image/png"
	"os"
)

type TwoFA struct {
	Config *Config
}

var twoFA *TwoFA

func Install2FA(configKeys ...string) *TwoFA {

	if twoFAConfig == nil {
		getTwoFAFromEnv(configKeys...)
	}

	return &TwoFA{twoFAConfig}
}

func Get2FAInstance() *TwoFA {
	if twoFA == nil {
		twoFA = Install2FA()
	}

	return twoFA
}

func (twoFA *TwoFA) GenerateOTP(accountName string) (*otp.Key, error) {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      twoFA.Config.Issuer,
		AccountName: accountName,
		Period:      twoFA.Config.Period,
		SecretSize:  twoFA.Config.SecretSize,
		Algorithm:   twoFA.Config.Algorithm,
	})

	if err != nil {
		return nil, err
	}

	return key, nil
}

func (twoFA *TwoFA) VerifyOTP(otp, secretKey string) bool {
	return totp.Validate(otp, secretKey)
}

func Test() {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "auction-dev",
		AccountName: "tiennc@gtv.vn",
		SecretSize:  15,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})

	image, err := key.Image(53,53)
	fmt.Println(err)
	fmt.Println(key.URL())

	f, err := os.Create("image.png")
	if err != nil {
		fmt.Println(err)
	}

	if err = png.Encode(f, image); err != nil {
		f.Close()
		fmt.Println(err)
	}

	if err = f.Close(); err != nil {
		fmt.Println(err)
	}




	//bar ,_ := qr.Encode(key.URL(),qr.L, qr.Auto)
	//bar, _ = barcode.Scale(bar, 30, 30)


	//totp.GenerateCode()
	//
	//totp.GenerateCodeCustom()
	//
	//totp.
	//
	//totp.

	//fmt.Println(totp.ValidateCustom("311180", "PQLZEKSB7O62ZSQUMSJYZVNN", time.Now(), totp.ValidateOpts{}))
	//fmt.Println(totp.Validate("735098", "PQLZEKSB7O62ZSQUMSJYZVNN"))
	//fmt.Println(totp.Validate("357863", "PQLZEKSB7O62ZSQUMSJYZVNN"))

	// otpauth://totp/auction-dev:tiennc@gtv.vn?secret=PQLZEKSB7O62ZSQUMSJYZVNN

	//fmt.Println(key.URL())
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(key)

}
