package elastic_search

import (
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v7"
	"sync"
)

type ElasticClient struct {
	client *elastic.Client
}

type ElasticClientManager struct {
	elasticClients sync.Map
}

var elasticClients = &ElasticClientManager{}

func InstallElasticClientManager(configKeys ...string) {

	getConfigFromEnv(configKeys...)

	if esConf == nil {
		err := fmt.Errorf("need config for elastic client first")
		panic(err)
	}

	client := NewElasticClient()
	if client == nil {
		err := fmt.Errorf("InstallElasticClientManager - NewElasticClient error")
		panic(err)
	}

	if val, ok := elasticClients.elasticClients.Load(esConf.Name); ok {
		fmt.Println(fmt.Sprintf("InstallElasticClientManager - config error: duplicated config.Name %s with %v", esConf.Name, val))

		return
	}

	elasticClients.elasticClients.Store(esConf.Name, &ElasticClient{client: client})
}

func GetElasticClient(dbName string) (client *ElasticClient) {
	if val, ok := elasticClients.elasticClients.Load(dbName); ok {
		if client, ok = val.(*ElasticClient); ok {
			return client
		}

	}

	panic(fmt.Sprintf("GetElasticClient - Not found client: %s", dbName))
}
