/* !!
 * File: producer.go
 * File Created: Wednesday, 5th May 2021 10:33:38 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 5th May 2021 10:35:00 am
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

package kafka_client

import (
	"fmt"
	"time"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/base"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/env"
	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"

	"github.com/Shopify/sarama"
)

// ClientKafka type;
type ClientKafka struct {
	config            *ConfigKafka
	producer          sarama.SyncProducer
	kafkaClientConfig *sarama.Config
}

var clientKafka *ClientKafka

// default value env key is "Kafka";
// if configKeys was set, key env will be first value (not empty) of this;
func InstallKafkaClient(configKeys ...string) *ClientKafka {
	if clientKafka != nil {
		return clientKafka
	}

	getKafkaConfigFromEnv(configKeys...)

	if kafkaConfig == nil {
		err := fmt.Errorf("need config for kafka client first")
		panic(err)
	}

	conf := sarama.NewConfig()
	conf.Consumer.Return.Errors = true
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Retry.Max = 5
	conf.Producer.Return.Successes = true
	conf.Producer.Compression = sarama.CompressionSnappy
	conf.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	conf.Producer.Flush.Frequency = 5 * time.Millisecond
	conf.ClientID = env.PodName

	// version, err := sarama.ParseKafkaVersion("2.1.1")
	// if err != nil {
	//  g_log.V(1).WithError(err).Errorf("ClientKafka::InstallKafkaClient - Error parsing Kafka version: %+v", err)

	// send log to telegram

	// 	return nil
	// }

	// conf.Version = version

	producer, err := sarama.NewSyncProducer(kafkaConfig.Addrs, conf)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("InstallKafkaClient - Error: %+v", err)

		// send log to telegram

		return nil
	}

	clientKafka = &ClientKafka{
		kafkaClientConfig: conf,
		config:            kafkaConfig,
		producer:          producer,
	}

	go clientKafka.CreateTopic()

	return clientKafka
}

// GetKafkaClientInstance func;
func GetKafkaClientInstance() *ClientKafka {
	if clientKafka == nil {
		return InstallKafkaClient()
	}

	return clientKafka
}

// Check exist topic
func checkExistTopic(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// CreateTopic func;
func (c *ClientKafka) CreateTopic() {
	if c == nil || c.config == nil {
		g_log.V(1).Error("ClientKafka::CreateTopic - Need InstallKafkaClient first")
		return
	}

	if len(c.config.ProducerTopics) == 0 {
		g_log.V(1).Error("ClientKafka::CreateTopic - Need least one topic")

		return
	}

	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = c.config.NumPartitions
	topicDetail.ReplicationFactor = c.config.ReplicationFactor
	topicDetail.ConfigEntries = make(map[string]*string)

	topicDetails := make(map[string]*sarama.TopicDetail)

	consumer, err := sarama.NewConsumer(c.config.Addrs, c.kafkaClientConfig)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("ClientKafka::CreateTopic - Error when create new consumer %+v", err)

		return
	}

	listTopic, _ := consumer.Topics()
	g_log.V(3).Infof("get topics available in kafka: %s", base.JSONDebugDataString(listTopic))

	listTopicNotExisted := make([]string, 0)
	for _, topicName := range c.config.ProducerTopics {
		if _, found := checkExistTopic(listTopic, topicName); !found {
			listTopicNotExisted = append(listTopicNotExisted, topicName)
		}
	}

	if len(listTopicNotExisted) == 0 {
		g_log.V(3).Infof("ClientKafka::CreateTopic - All of %+v was existed", c.config.ProducerTopics)

		return
	}

	for _, topicName := range listTopicNotExisted {
		topicDetails[topicName] = topicDetail
	}

	request := sarama.CreateTopicsRequest{
		Timeout:      time.Second * 15,
		TopicDetails: topicDetails,
	}

	// Send request to Broker
	broker := sarama.NewBroker(c.config.Addrs[0])
	broker.Open(c.kafkaClientConfig)

	response, err := broker.CreateTopics(&request)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("ClientKafka::CreateTopic - CreateTopics Error: %+v", err)

		return
	}

	t := response.TopicErrors
	for key, val := range t {
		if val.Err != sarama.ErrNoError {
			g_log.V(1).WithError(err).Errorf("😡😡😡 ClientKafka::CreateTopic - Create topic key: %s - Error: %s at pod %s in host %s", key, val.Err.Error(), env.PodName, env.HostName)
		}
	}
}

// ProducerPushMessage func;
func (c *ClientKafka) ProducerPushMessage(topic string, messageObj MessageKafka) (partition int32, offset int64, err error) {
	// debug.PrintStack()
	if c == nil || c.producer == nil {
		return 0, 0, fmt.Errorf("ClientKafka::ProducerPushMessage - Not found any producer")
	}

	g_log.V(4).Infof("ClientKafka::ProducerPushMessage - Push to topic: %s object data: %s", topic, base.JSONDebugDataString(messageObj))

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(base.JSONDebugDataString(messageObj)),
	}

	return c.producer.SendMessage(msg)
}

// ProducerPushMessage func;
func (c *ClientKafka) ProducerPushMessageWithKey(topic, key string, messageObj MessageKafka) (partition int32, offset int64, err error) {
	// debug.PrintStack()
	if c == nil || c.producer == nil {
		return 0, 0, fmt.Errorf("ClientKafka::ProducerPushMessage - Not found any producer")
	}

	g_log.V(4).Infof("ClientKafka::ProducerPushMessage - Push to topic: %s, key: %s, object data: %s", topic, key, base.JSONDebugDataString(messageObj))

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(base.JSONDebugDataString(messageObj)),
	}

	return c.producer.SendMessage(msg)
}
