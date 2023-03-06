/* !!
 * File: main.go
 * File Created: Wednesday, 5th May 2021 10:33:38 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 5th May 2021 10:34:33 am
 * Modified By: KimEricko™ (phamkim.pr@gmail.com>)
 * -----
 * Copyright 2018 - 2021 GTV, GGroup
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Developer: NhokCrazy199 (phamkim.pr@gmail.com)
 */

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/kafka_client"
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
