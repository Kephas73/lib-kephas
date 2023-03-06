/* !!
 * File: main.go
 * File Created: Wednesday, 5th May 2021 10:33:38 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 5th May 2021 10:34:39 am
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
	"time"

	"github.com/spf13/viper"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/kafka_client"
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
