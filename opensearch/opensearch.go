package opensearch

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "github.com/Kephas73/lib-kephas/base"
    "github.com/Kephas73/lib-kephas/document"
    "github.com/opensearch-project/opensearch-go"
    "github.com/opensearch-project/opensearch-go/opensearchapi"
    "io"
    "io/ioutil"
    "net/http"
    "time"
)

type OpenSearch struct {
    Hostname    []string
    Username    string
    Password    string
    Connection  *opensearch.Client
    Timeout     int
    IndexFormat string
}

func New(host []string, username, password, index string, timeout int) *OpenSearch {
    l := OpenSearch{}
    l.Hostname = host
    l.Username = username
    l.Password = password
    l.Connection = nil
    l.Timeout = timeout
    l.IndexFormat = index
    return &l
}

func (l *OpenSearch) Connect() (*opensearch.Client, error) {
    client, err := opensearch.NewClient(opensearch.Config{
        Transport: &http.Transport{
            MaxIdleConns:        50,
            MaxIdleConnsPerHost: 50,
        },
        Addresses: l.Hostname,
        Username:  l.Username, // For testing only. Don't store credentials in code.
        Password:  l.Password,
    })

    if err != nil {
        return client, err
    }

    if client != nil {
        l.Connection = client
    }

    return client, nil
}

func (l *OpenSearch) CreateIndex(index string, mapping interface{}) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(l.Timeout)*time.Second)
    defer cancel()

    req := opensearchapi.IndicesCreateRequest{
        Index: index,
        Body:  bytes.NewReader(base.JSONDebugData(mapping)),
    }

    if _, err := req.Do(ctx, l.Connection); err != nil {
        return err
    }
    return nil
}

func (l *OpenSearch) InsertDocument(index, id string, object interface{}) error {
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(l.Timeout)*time.Second)
    defer cancel()

    req := opensearchapi.IndexRequest{
        Index:      index,
        DocumentID: id,
        Body:       bytes.NewReader(base.JSONDebugData(object)),
    }

    if _, err := req.Do(ctx, l.Connection); err != nil {
        return err
    }
    return nil
}

func (l *OpenSearch) CountDocument(index []string, bodyQuery io.Reader) (document.ResponseOpenSearch, error) {
    var result document.ResponseOpenSearch
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(l.Timeout)*time.Second)
    defer cancel()

    listQueries := make([]func(*opensearchapi.CountRequest), 0)
    listQueries = append(listQueries, l.Connection.Count.WithIndex(index...))
    listQueries = append(listQueries, l.Connection.Count.WithBody(bodyQuery))
    listQueries = append(listQueries, l.Connection.Count.WithContext(ctx))

    res, err := l.Connection.Count(listQueries...)
    if err != nil {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Count request error: %+v", err))
        return result, err
    }

    dataLog, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Can not parse the body of response error: %+v", err))

        return result, err
    }

    if res.IsError() {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Count request is error: %s", dataLog))

        return result, err
    }

    if err = json.Unmarshal(dataLog, &result); err != nil {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Can not parse the response error: %+v", err))

        return result, err
    }

    return result, nil
}

func (l *OpenSearch) SearchDocument(index []string, bodyQuery io.Reader) (document.ResponseOpenSearch, error) {
    var result document.ResponseOpenSearch
    ctx, cancel := context.WithTimeout(context.Background(), time.Duration(l.Timeout)*time.Second)
    defer cancel()

    listQueries := make([]func(request *opensearchapi.SearchRequest), 0)
    listQueries = append(listQueries, l.Connection.Search.WithIndex(index...))
    listQueries = append(listQueries, l.Connection.Search.WithBody(bodyQuery))
    listQueries = append(listQueries, l.Connection.Search.WithContext(ctx))

    res, err := l.Connection.Search(listQueries...)
    if err != nil {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Count request error: %+v", err))
        return result, err
    }

    dataLog, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Can not parse the body of response error: %+v", err))

        return result, err
    }

    if res.IsError() {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Count request is error: %s", dataLog))
        return result, fmt.Errorf(string(dataLog))
    }

    if err = json.Unmarshal(dataLog, &result); err != nil {
        fmt.Println(fmt.Sprintf("OpenSearch::CountDocument - Can not parse the response error: %+v", err))
        return result, err
    }

    return result, nil
}

func (l *OpenSearch) IndexDefault() string {
    return fmt.Sprintf("%s-%s", l.IndexFormat, time.Now().Format("2006.01.02"))
}
