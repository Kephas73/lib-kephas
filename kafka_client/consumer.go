/* !!
 * File: consumer.go
 * File Created: Wednesday, 5th May 2021 10:33:38 am
 * Author: KimEricko™ (phamkim.pr@gmail.com)
 * -----
 * Last Modified: Wednesday, 5th May 2021 10:34:56 am
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
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"gitlab.gplay.vn/gtv-backend/fountain/baselib/g_log"

	"github.com/Shopify/sarama"
)

// ConsumerProcessInstance func;
type ConsumerProcessInstance interface {
	ErrorCallback(err error)
	MessageCallback(messageObj MessageKafka)
}

// InstallConsumerGroup func;
// Func need run as goroutine;
// If groupName is not set, groupName will be env ConsumerTopicNames
func (c *ClientKafka) InstallConsumerGroup(processingInstance ConsumerProcessInstance, consumerGroupName string, topicName ...string) {
	g_log.V(1).Infof("ClientKafka::InstallConsumerGroup - Register consumer group: %s for topics: %+v", consumerGroupName, topicName)

	if c == nil || c.config == nil {
		g_log.V(1).Error("ClientKafka::InstallConsumerGroup - Need InstallKafkaClient first")

		// send to telegram

		return
	}

	if len(topicName) == 0 {
		err := fmt.Errorf("need a least 1 topic name")
		g_log.V(1).WithError(err).Errorf("ClientKafka::InstallConsumerGroup - InstallKafkaClient error: %+v", err)
		panic(err)
	}

	c.kafkaClientConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	c.kafkaClientConfig.Consumer.Return.Errors = true

	// c.kafkaClientConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	c.kafkaClientConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	// c.kafkaClientConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange

	consumer := newConsumerInstance(processingInstance)

	ctx, cancel := context.WithCancel(context.Background())

	client, err := sarama.NewConsumerGroup(c.config.Addrs, consumerGroupName, c.kafkaClientConfig)
	if err != nil {
		g_log.V(1).WithError(err).Errorf("ClientKafka::InstallConsumerGroup - Error creating consumer group client: %+v", err)

		// Send to telegram

		cancel()

		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// `Consume` should be called inside an infinite loop, when a
				// server-side rebalance happens, the consumer session will need to be
				// recreated to get the new claims
				if err := client.Consume(ctx, topicName, consumer); err != nil {
					processingInstance.ErrorCallback(err)
				}
				// check if context was cancelled, signaling that the consumer should stop
				if ctx.Err() != nil {
					return
				}
				consumer.ready = make(chan bool, 1)
			}

		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case err := <-client.Errors():
				processingInstance.ErrorCallback(err)
			case <-ctx.Done():
				return
			}
		}
	}()

	topicsStr := strings.Join(topicName, ", ")

	<-consumer.ready // Await till the consumer has been set up
	g_log.V(1).Infof("ClientKafka::InstallConsumerGroup - Consumer group: %q for topics: %q up and running!...", consumerGroupName, topicsStr)

	<-ctx.Done()
	g_log.V(1).Infof("ClientKafka::InstallConsumerGroup - Consumer group: %q for topics: %q. Terminating: context cancelled", consumerGroupName, topicsStr)

	cancel()

	wg.Wait()
	if err = client.Close(); err != nil {
		g_log.V(1).WithError(err).Errorf("ClientKafka::InstallConsumerGroup - Consumer group: %s for topics: %s. Error closing client: %v", consumerGroupName, topicsStr, err)
	}
}

// Consumer represents a Sarama consumer group consumer
type consumerInstance struct {
	ready              chan bool
	processingInstance ConsumerProcessInstance
}

// NewConsumer func;
func newConsumerInstance(processingInstance ConsumerProcessInstance) *consumerInstance {
	// Mark the consumer as ready
	return &consumerInstance{
		ready:              make(chan bool),
		processingInstance: processingInstance,
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *consumerInstance) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *consumerInstance) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *consumerInstance) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29

	for message := range claim.Messages() {
		msgObj := MessageKafka{}
		err := json.Unmarshal(message.Value, &msgObj)
		if err != nil {
			consumer.processingInstance.ErrorCallback(err)
		} else {
			msgObj.Topic = message.Topic
			consumer.processingInstance.MessageCallback(msgObj)
		}

		session.MarkMessage(message, "")
	}

	return nil
}
