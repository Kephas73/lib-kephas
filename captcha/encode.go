package captcha

import (
	"encoding/csv"
	"github.com/Kephas73/lib-kephas/constant"
	"github.com/jszwec/csvutil"
	"io"
	"math/rand"
	"os"
	"time"
	"unicode/utf8"
)

type Captcha struct {
	CaptchaID    string `json:"captcha_id,omitempty" csv:"captcha_id"`
	CaptchaValue string `json:"captcha_value,omitempty" csv:"captcha_value"`
	Image        string `json:"image,omitempty" csv:"image"`
}

// NewCaptcha func:
// Only accept csv. files, default format, no header
func NewCaptcha(pathKey string) (captcha []*Captcha) {
	if pathKey == constant.StrEmpty {
		return nil
	}

	file, err := os.Open(pathKey)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)
	csvReader.TrimLeadingSpace = constant.ModeTrimLeadingSpace
	csvReader.LazyQuotes = constant.ModeLazyQuotes
	csvReader.FieldsPerRecord = constant.ModeFieldsPerRecord
	comma, _ := utf8.DecodeRuneInString(constant.ModeComma)
	csvReader.Comma = comma

	header, err := csvutil.Header(Captcha{}, "csv")
	if err != nil {
		panic(err)
	}

	dec, err := csvutil.NewDecoder(csvReader, header...)
	if err != nil {
		panic(err)
	}

	for {
		var u Captcha
		if err = dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		captcha = append(captcha, &u)
	}

	return
}

func (captcha *CaptchaClient) RandomCaptcha() *Captcha {

	if len(captcha.captcha) == constant.ValueEmpty || captcha.captcha == nil {
		return nil
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return captcha.captcha[r.Intn(len(captcha.captcha))]
}
