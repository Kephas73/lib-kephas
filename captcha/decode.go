package captcha

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var tr = &http.Transport{
	MaxIdleConns:          50,
	MaxIdleConnsPerHost:   50,
	ResponseHeaderTimeout: 30 * time.Second,
}

type Task struct {
	ClientKey string `json:"clientKey"`
	Task      struct {
		Type      string `json:"type"`
		Body      string `json:"body"`
		Phrase    bool   `json:"phrase"`
		Case      bool   `json:"case"`
		Numeric   int    `json:"numeric"`
		Math      bool   `json:"math"`
		MinLength int    `json:"minLength"`
		MaxLength int    `json:"maxLength"`
	} `json:"task"`
	SoftId  int   `json:"softId"`
	TaskId  int64 `json:"taskId"`
	ErrorId int   `json:"errorId"`
}

type TaskResult struct {
	ErrorId  int    `json:"errorId"`
	Status   string `json:"status"`
	Solution struct {
		Text string `json:"text"`
		Url  string `json:"url"`
	} `json:"solution"`
	Cost       string `json:"cost"`
	IP         string `json:"ip"`
	CreateTime int32  `json:"createTime"`
	EndTime    int32  `json:"endTime"`
	SolveCount int    `json:"solveCount"`
}

func CreateTask(image64 string) (*Task, error) {
	body := Task{}
	body.ClientKey = configCaptcha.Decode.ClientKey
	body.Task.Type = "ImageToTextTask"
	body.Task.Body = image64
	body.Task.Phrase = false
	body.Task.Case = false
	body.Task.Numeric = 0
	body.Task.Math = false
	body.Task.MinLength = 0
	body.Task.MaxLength = 0
	body.SoftId = 0

	dataByte, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", configCaptcha.Decode.Url.Create, bytes.NewBuffer(dataByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout:   time.Second * 30,
		Transport: tr,
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	resData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resObj := &Task{}
	if err = json.Unmarshal(resData, resObj); err != nil {
		return nil, err
	}

	return resObj, nil
}

func GetTask(taskId int64) (*TaskResult, error) {
	body := Task{}
	body.ClientKey = configCaptcha.Decode.ClientKey
	body.TaskId = taskId

	dataByte, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", configCaptcha.Decode.Url.Get, bytes.NewBuffer(dataByte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout:   time.Second * 30,
		Transport: tr,
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("not 200")
	}

	resData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	resObj := &TaskResult{}
	if err = json.Unmarshal(resData, resObj); err != nil {
		return nil, err
	}

	return resObj, nil
}

func (captcha *CaptchaClient) ImageToText(image64 string) (*TaskResult, error) {
	task, err := CreateTask(image64)
	if err != nil {
		return nil, err
	}

	// Check if the result is ready, if not loop until it is
	response, err := GetTask(task.TaskId)
	if err != nil {
		return nil, err
	}

	for {
		if response.Status == "processing" {
			time.Sleep(2 * time.Second)
			response, err = GetTask(task.TaskId)
			if err != nil {
				return nil, err
			}
		} else {
			break
		}
	}

	return response, nil
}
