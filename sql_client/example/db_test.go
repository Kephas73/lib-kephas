package example

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/env"
	"github.com/Kephas73/lib-kephas/sql_client"
	"testing"
)

func init() {
	env.SetupConfigEnv("config.json")
}

func TestDB(t *testing.T)  {

	sql_client.InstallSQLClientManager()
	err := sql_client.GetSQLClient("broker-dev").Get().Ping()
	fmt.Println(err)
}