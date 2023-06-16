package opensearch

import (
    "errors"
    "fmt"
    "github.com/Kephas73/go-lib/document"
    "github.com/Kephas73/go-lib/logger"
    "github.com/Kephas73/go-lib/util"
    "time"
)

type OpenSearchClient struct {
    clients *OpenSearch
    index   func() string
}

var (
    openSearchClient *OpenSearchClient
)

func InstallOpenSearchClient() *OpenSearchClient {
    createConfigFromEnv()

    if openSearchClient == nil {
        mC := New(openSearchConf.Hosts, openSearchConf.Username, openSearchConf.Password, openSearchConf.IndexFormat, openSearchConf.Timeout)
        conn, err := mC.Connect()
        if err != nil {
            err := fmt.Errorf("InstallOpenSearchClient - Can not create connection to %s - Error: %+v", mC.Hostname, err)
            logger.Error("InstallOpenSearchClient - Error: %+v", err)

            panic(err)
        }

        mC.Connection = conn
        openSearchClient = &OpenSearchClient{
            clients: mC,
            index:   mC.IndexDefault,
        }
        logger.Info("Connect(): conn: %+v ", conn)
    }

    return openSearchClient
}

func GetOpenSearchClient() *OpenSearchClient {
    if openSearchClient == nil {
        openSearchClient = InstallOpenSearchClient()
    }

    return openSearchClient
}

func (o *OpenSearchClient) InsertDocument(document document.Document) error {
    logger.Info("OpenSearchClient::Insert - OpenSearchClient Addr: %p - With list conn: %+v", o, o.clients)

    return o.insert(document)
}

func (o *OpenSearchClient) insert(document document.Document) error {
    logger.Info("OpenSearchClient::insert - openSearchClient Addr: %p - With list conn: %+v", o, o.clients)

    document.TimeStamp = util.GetTimestampData()

    retries := 0
    for retries < 3 {
        err := o.clients.InsertDocument(o.index(), document.IDDoc, document)
        if err != nil {
            logger.Error("OpenSearchClient::insert - Write to %v - err: %+v", o.clients.Connection, err)
            retries++
            if retries > 3 {
                return err
            }
            time.Sleep(time.Second)
        } else {
            break
        }
    }

    if retries >= 3 {
        logger.Info("OpenSearchClient::insert - Retries too much")
        return errors.New("retries too much")
    }
    return nil
}
