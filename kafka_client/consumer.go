package kafka_client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

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
	fmt.Printf(fmt.Sprintf("ClientKafka::InstallConsumerGroup - Register consumer group: %s for topics: %v", consumerGroupName, topicName))

	if c == nil || c.config == nil {
		fmt.Println("ClientKafka::InstallConsumerGroup - Need InstallKafkaClient first")

		// send to telegram

		return
	}

	if len(topicName) == 0 {
		err := fmt.Errorf("need a least 1 topic name")
		fmt.Printf(fmt.Sprintf("ClientKafka::InstallConsumerGroup - InstallKafkaClient error: %+v", err))
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
		fmt.Printf(fmt.Sprintf("ClientKafka::InstallConsumerGroup - Error creating consumer group client: %v", err))

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
	fmt.Printf(fmt.Sprintf("ClientKafka::InstallConsumerGroup - Consumer group: %q for topics: %q up and running!...", consumerGroupName, topicsStr))

	<-ctx.Done()
	fmt.Printf(fmt.Sprintf("ClientKafka::InstallConsumerGroup - Consumer group: %q for topics: %q. Terminating: context cancelled", consumerGroupName, topicsStr))

	cancel()

	wg.Wait()
	if err = client.Close(); err != nil {
		fmt.Printf(fmt.Sprintf("ClientKafka::InstallConsumerGroup - Consumer group: %s for topics: %s. Error closing client: %v", consumerGroupName, topicsStr, err))
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
