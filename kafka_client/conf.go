package kafka_client

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// ConfigKafka type;
type ConfigKafka struct {
	Addrs             []string `json:"addrs,omitempty"`
	ReplicationFactor int16    `json:"replication_factor,omitempty"`
	NumPartitions     int32    `json:"num_partitions,omitempty"`
	ProducerTopics    []string `json:"producer_topic,omitempty"`
}

// MessageKafka func;
type MessageKafka struct {
	Topic      string `json:"topic"`
	Event      string `json:"event"`
	ObjectJSON string `json:"object_json"`
}

// LogMessageKafka type
type LogMessageKafka struct {
	ServerName     string `json:"server_name"`
	FileName       string `json:"file_name"`
	Line           int32  `json:"line"`
	TimeStamp      string `json:"time_stamp"`
	TimeAccessHash int64  `json:"time_access_hash"`
	LogType        string `json:"log_type"`
	LogMessage     string `json:"log_message"`
}

const (
	defaultNumPartitions     = int32(1)
	defaultReplicationFactor = int16(1)
)

var (
	kafkaConfig *ConfigKafka
)

// default value env key is "Kafka";
// if configKeys was set, key env will be first value (not empty) of this;
func getKafkaConfigFromEnv(configKeys ...string) {
	configKey := "Kafka"
	for _, envKey := range configKeys {
		envKeyTrim := strings.TrimSpace(envKey)
		if envKeyTrim != "" {
			configKey = envKeyTrim
		}
	}

	kafkaConfig = &ConfigKafka{}

	if err := viper.UnmarshalKey(configKey, kafkaConfig); err != nil {
		err := fmt.Errorf("not found config with env %q for kafka with error: %+v", configKey, err)
		panic(err)
	}

	if len(kafkaConfig.Addrs) == 0 {
		err := fmt.Errorf("not found any addr as host for kafka at %q", fmt.Sprintf("%s.Addrs", configKey))
		panic(err)
	}

	if kafkaConfig.NumPartitions == 0 {
		kafkaConfig.NumPartitions = defaultNumPartitions
	}

	if kafkaConfig.ReplicationFactor == 0 {
		kafkaConfig.ReplicationFactor = defaultReplicationFactor
	}

	if len(kafkaConfig.ProducerTopics) == 0 {
		err := fmt.Errorf("not found producer topic config with env %q for kafka", fmt.Sprintf("%s.ProducerTopics", configKey))
		panic(err)
	}
}
