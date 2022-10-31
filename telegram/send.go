package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type RawData struct {
	ChatId    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

type ResponseTele struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

type Telegram struct {
	Config *Config
}

var (
	t  *Telegram
	tr = &http.Transport{
		MaxIdleConns:          50,
		MaxIdleConnsPerHost:   50,
		ResponseHeaderTimeout: 30 * time.Second,
	}
)

func InstallTelegram(configKeys ...string) *Telegram {

	if teleConfig == nil {
		getTeleConfigFromEnv(configKeys...)
	}

	return &Telegram{teleConfig}
}

func GetTeleInstance() *Telegram {
	if t == nil {
		t = InstallTelegram()
	}

	return t
}

func (t *Telegram) SendMessage(raw RawData) error {

	teleEndpoint := fmt.Sprintf("%s/bot%s/sendMessage", t.Config.Bot.Endpoint, t.Config.Bot.Token)
	if raw.ParseMode == "" {
		raw.ParseMode = t.Config.Bot.ParseMode
	}

	jsonReq, err := json.Marshal(raw)
	if err != nil {
		return err
	}

	rq, err := http.NewRequest("POST", teleEndpoint, bytes.NewBuffer(jsonReq))
	if err != nil {
		return err
	}
	rq.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout:   time.Second * time.Duration(teleConfig.Bot.Timeout),
		Transport: tr,
	}

	res, err := client.Do(rq)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		result := ResponseTele{}
		data, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		err = json.Unmarshal(data, &result)
		if err != nil {
			return err
		}

		return errors.New(result.Description)
	}

	return nil
}