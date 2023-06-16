package elastic_search

import (
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/spf13/viper"
	"strings"
)

type ElasticConfig struct {
	Name        string
	Environment string
	Hosts       []string
	Username    string
	Password    string
}

var esConf *ElasticConfig

func getConfigFromEnv(configKeys ...string) {
	configKey := "Elastic"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	esConf = &ElasticConfig{}

	if err := viper.UnmarshalKey(configKey, esConf); err != nil {
		err := fmt.Errorf("not found config with env %q for Elastic with error: %+v", configKey, err)
		panic(err)
	}

	if esConf.Name == "" {
		err := fmt.Errorf("not found config name with env %q for Elastic", fmt.Sprintf("%s.Name", configKey))
		panic(err)
	}

	if esConf.Environment == "" {
		err := fmt.Errorf("not found config environment with env %q for Elastic", fmt.Sprintf("%s.Environment", configKey))
		panic(err)
	}

	if len(esConf.Hosts) == 0 {
		err := fmt.Errorf("not found config hosts with env %q for Elastic", fmt.Sprintf("%s.Hosts", configKey))
		panic(err)
	}
}

func NewElasticClient() (client *elastic.Client) {
	cfg := elastic.Config{
		Addresses: esConf.Hosts,
		Username:  esConf.Username,
		Password:  esConf.Password,
	}
	var err error

	client, err = elastic.NewClient(cfg)
	if err != nil {
		fmt.Println("NewElasticClient - Can't connect to Elastic Search server...!")
		panic(err)
	}

	fmt.Println("NewElasticClient - ElasticManager initialized successfully!")

	return client
}