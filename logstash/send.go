package logstash

import (
    "errors"
    "fmt"
    "github.com/Kephas73/lib-kephas/base"
    "github.com/Kephas73/lib-kephas/document"
    "github.com/Kephas73/lib-kephas/logger"
    "math"
    "math/rand"
    "time"
)

type LogStashClient struct {
    clients  []*Logstash
    curIndex int
}

var (
    logStashClient *LogStashClient
)

// InstallLogStashClient func
func InstallLogStashClient() *LogStashClient {
    createConfigFromEnv()

    count := 0
    for logStashClient == nil && count < 5 {
        count++
        myG := NewList(logStashConf.Hosts, int(math.MaxInt32))
        for _, g := range myG {
            conn, err := g.Connect()
            if err != nil {
                err := fmt.Errorf("InstallLogStashClient - Can not create connection to %s:%d - Error: %+v", g.Hostname, g.Port, err)
                fmt.Println(fmt.Sprintf("Gostash::InstallLogStashClient - Error: %+v", err))

                panic(err)
            }

            g.Connection = conn
            fmt.Println(fmt.Sprintf("Connect(): conn: %+v ", conn))
        }

        logStashClient = &LogStashClient{
            clients:  myG,
            curIndex: 0,
        }

        time.Sleep(500 * time.Millisecond)
    }

    if count >= 5 {
        err := fmt.Errorf("installLogStashClient - Can not create logstash's instace")
        panic(err)
    }

    return logStashClient
}

// GetLogStashClient func
func GetLogStashClient() *LogStashClient {
    if logStashClient == nil {
        logStashClient = InstallLogStashClient()
    }

    return logStashClient
}

// GetNextClient func;
func (e *LogStashClient) GetNextClient() int {
    e.curIndex++
    if e.curIndex >= len(e.clients) {
        e.curIndex = 0
    }

    return e.curIndex
}

///////////////////////////////////////////
// Cho phép nhiều con chạy
func random() int {
    return rand.Intn(1)
}

func (e *LogStashClient) InsertDocument(document document.Document) error {
    fmt.Println(fmt.Sprintf("LogStashClient::InsertDocument - logStashClient Addr: %p - With list conn: %+v", e, e.clients))

    return e.insert(document)
}

func (e *LogStashClient) insert(document document.Document) error {
    fmt.Println(fmt.Sprintf("LogStashClient::insert - logStashClient Addr: %p - With list conn: %+v", e, e.clients))

    document.TimeStamp = base.GetTimestampData()

    retries := 0
    for retries < 3 {
        if len(e.clients) == 0 {
            logger.Info("LogStashClient::insert - No connection!")
            time.Sleep(1 * time.Second)
            retries++

            continue
        }

        rand.Seed(time.Now().UnixNano())
        myIdx := random()
        // myIdx := e.GetNextClient()
        myConn := e.clients[myIdx]
        err := myConn.Writeln(base.JSONDebugDataString(document))

        if err != nil {
            fmt.Println(fmt.Sprintf("LogStashClient::insert - Write to %v - err: %+v", myConn.Connection, err))
            fmt.Println(fmt.Sprintf("LogStashClient::insert - InsertLog into [%s] was error: %+v", myConn.String(), err))
            retries++
            if retries > 3 {
                return err
            }

            time.Sleep(1 * time.Second)
        } else {
            break
        }
    }

    if retries >= 3 {
        fmt.Println(fmt.Sprintf("LogStashClient::insert - Retries too much"))
        return errors.New("retries too much")
    }
    return nil
}
