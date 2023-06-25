package document

import (
    "encoding/json"
    "github.com/Kephas73/lib-kephas/base"
    "time"
)

type Document struct {
    TimeStamp string   `json:"@timestamp"`
    IDDoc     string   `json:"id_doc"`
    Document  DataBase `json:"document"`
}

func MakeDocument() Document {
    return Document{
        Document: NewDefaultData(),
    }
}

// ToString func
func (e *Document) ToString() string {
    return base.JSONDebugDataString(e)
}

func (e *Document) RandomIDDoc() {
    e.IDDoc = base.RandStringRunes(21) // 21: vì thấy trong opensearch hiện tại đang để 21
}

func (e *Document) SetIDDoc(id string) {
    e.IDDoc = id
}

type DataBase interface {
    FromJSON(data []byte) error
    ToJSON() string
    SetEventName(eventName string) DataBase
    SeDataJSON(dataJSON interface{}) DataBase
    SetDescription(description string) DataBase
}

type DefaultData struct {
    EventName    string      `json:"event_name,omitempty"`
    Data         interface{} `json:"data,omitempty"`
    Description  string      `json:"description,omitempty"`
    TimeStarted  int64       `json:"time_started,omitempty"`
    TimeFinished int64       `json:"time_finished,omitempty"`
    TimeExecute  int64       `json:"time_execute,omitempty"`
}

func (obj *DefaultData) FromJSON(data []byte) error {
    return json.Unmarshal(data, obj)
}

func (obj *DefaultData) ToJSON() string {
    return base.JSONDebugDataString(obj)
}

func (obj *DefaultData) SetEventName(eventName string) DataBase {
    obj.EventName = eventName

    return obj
}

func (obj *DefaultData)  SeDataJSON(dataJSON interface{}) DataBase {
    obj.Data = dataJSON

    return obj
}

func (obj *DefaultData) SetDescription(description string) DataBase {
    obj.Description = description

    return obj
}

func NewDefaultData() *DefaultData {
    timeNow := time.Now().Unix()
    return &DefaultData{
        EventName:    "Default",
        TimeStarted:  timeNow,
        TimeFinished: timeNow,
        TimeExecute:  0,
    }
}
