package main

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/kafka_client"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.Set("Kafka.Broker", "kafka-chatting.gtvplusdev.info.private:9092")
	viper.Set("Kafka.ProducerTopics", "gtv_test")
	viper.Set("Kafka.ConsumerGroupName", "gtv_consumer_test")
	viper.Set("Kafka.ConsumerTopicNames", "gtv_test")
}

func main() {
	kafka_client.InstallKafkaClient()

	kafClient := kafka_client.GetKafkaClientInstance()

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("This is object: %d", i)
		messageObj := kafka_client.MessageKafka{
			Event:      "Event1",
			ObjectJSON: msg,
		}
		par, off, err := kafClient.ProducerPushMessage("gtv_test", messageObj)
		fmt.Printf("par: %d, off: %d, msg: %s, err: %+v \n", par, off, msg, err)
		time.Sleep(2 * time.Second)
	}
}
