package main

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/env"
	"github.com/Kephas73/lib-kephas/kafka_client"

	"github.com/spf13/viper"
)

func init() {
	env.SetupConfigEnv("config.json")
	viper.Set("Kafka.Broker", "18.142.112.173:9092")
	viper.Set("Kafka.ProducerTopics", "auction_test")
	viper.Set("Kafka.ConsumerGroupName", "auction_consumer_test")
	viper.Set("Kafka.ConsumerTopicNames", "auction_test")
}

func main() {
	kafka_client.InstallKafkaClient()

	kafClient := kafka_client.GetKafkaClientInstance()

	consumerCallback := &ConsumerInstace{}
	kafClient.InstallConsumerGroup(consumerCallback, "gtv_test")

	fmt.Printf("Consumer done!\n")
}

// ConsumerInstace func;
type ConsumerInstace struct{}

// ErrorCallback func;
func (c *ConsumerInstace) ErrorCallback(err error) {
	fmt.Printf("errCallback::Error - %+v\n", err)
}

// MessageCallback func;
func (c *ConsumerInstace) MessageCallback(msgObj kafka_client.MessageKafka) {
	fmt.Printf("procCallback::msgObj - %s\n", base.JSONDebugDataString(msgObj))
}
