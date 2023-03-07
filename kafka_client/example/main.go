package main

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/kafka_client"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

func init() {
	viper.Set("Kafka.Broker", "kafka-chatting.gtvplusdev.info.private:9092")
	viper.Set("Kafka.ProducerTopics", "gtv_test")
	viper.Set("Kafka.ConsumerGroupName", "gtv_consumer_test")
	viper.Set("Kafka.ConsumerTopicNames", "gtv_test")
}

type DataTmp struct {
	Data int64 `json:"data,omitempty"`
}

func main() {
	instance := kafka_client.InstallKafkaClient()
	if instance == nil {
		err := fmt.Errorf("installKafkaClient - Can create instance for kafka")
		panic(err)
	}

	randInt := rand.Int63()
	tmp := DataTmp{
		Data: randInt,
	}

	dataJSON := base.JSONDebugDataString(tmp)
	msg := kafka_client.MessageKafka{
		Event:      "Test",
		ObjectJSON: dataJSON,
	}

	nullInstance := &KafkaListen{}

	go instance.InstallConsumerGroup(nullInstance, "gtv_test")

	go func() {
		for {
			par, off, err := instance.ProducerPushMessage("gtv_test", msg)
			if err != nil {
				err := fmt.Errorf("testProducerPushMessage - ProducerPushMessage Error %+v while result expect nil", err)
				panic(err)
			}

			if off == 0 {
				err := fmt.Errorf("testProducerPushMessage - ProducerPushMessage offset is 0 while result expect greater 0")
				panic(err)
			}

			_ = par

			time.Sleep(time.Second)
		}
	}()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Waiting")
	<-sigterm
	log.Println("Done!")
}

type KafkaListen struct{}

func (instance *KafkaListen) ErrorCallback(err error) {
	log.Printf("processingError: %+v\n", err)
}

func (instance *KafkaListen) MessageCallback(messageObj kafka_client.MessageKafka) {
	log.Printf("processingMessage: %+v\n", messageObj)
}
