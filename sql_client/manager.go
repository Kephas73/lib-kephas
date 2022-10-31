package sql_client

import (
	"fmt"
	"sync"
)

type sqlClientManager struct {
	sqlClients sync.Map
}

var sqlClientsManagerInstance = &sqlClientManager{}

func InstallSQLClientManager(configKeys ...string) {
	getConfigFromEnv(configKeys...)

	for _, config := range configs {
		client := NewSqlxDB(config)
		if client == nil {
			err := fmt.Errorf("InstallSQLClientsManager - NewSqlxDB {%v} error", config)
			panic(err)
		}

		if config.Name == "" {
			err := fmt.Errorf("InstallSQLClientsManager - config error: config.Name is empty")
			panic(err)
		}
		if val, ok := sqlClientsManagerInstance.sqlClients.Load(config.Name); ok {
			err := fmt.Errorf("InstallSQLClientsManager - config error: duplicated config.Name {%v}", val)
			panic(err)
		}

		sqlClientsManagerInstance.sqlClients.Store(config.Name, client)
	}
}

// GetSQLClient type;
func GetSQLClient(dbName string) (client *SQLClient) {
	if val, ok := sqlClientsManagerInstance.sqlClients.Load(dbName); ok {
		if client, ok = val.(*SQLClient); ok {
			return
		}
	}
	return
}

// GetSQLClientManager type;
func GetSQLClientManager() sync.Map {
	return sqlClientsManagerInstance.sqlClients
}