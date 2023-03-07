package kafka_client

import (
	"fmt"
	"github.com/Kephas73/lib-kephas/base"
	"github.com/Kephas73/lib-kephas/env"
	"time"

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
	conf.ClientID = fmt.Sprintf("%s_%d", env.Environment.Server.ServiceName, base.TimeNowUnix())

	// version, err := sarama.ParseKafkaVersion("2.1.1")
	// if err != nil {
	//  g_log.V(1).WithError(err).Errorf("ClientKafka::InstallKafkaClient - Error parsing Kafka version: %+v", err)

	// send log to telegram

	// 	return nil
	// }

	// conf.Version = version

	producer, err := sarama.NewSyncProducer(kafkaConfig.Addrs, conf)
	if err != nil {
		fmt.Println(fmt.Sprintf("InstallKafkaClient - Error: %+v", err))

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
		fmt.Println("ClientKafka::CreateTopic - Need InstallKafkaClient first")
		return
	}

	if len(c.config.ProducerTopics) == 0 {
		fmt.Println("ClientKafka::CreateTopic - Need least one topic")

		return
	}

	topicDetail := &sarama.TopicDetail{}
	topicDetail.NumPartitions = c.config.NumPartitions
	topicDetail.ReplicationFactor = c.config.ReplicationFactor
	topicDetail.ConfigEntries = make(map[string]*string)

	topicDetails := make(map[string]*sarama.TopicDetail)

	consumer, err := sarama.NewConsumer(c.config.Addrs, c.kafkaClientConfig)
	if err != nil {
		fmt.Printf(fmt.Sprintf("ClientKafka::CreateTopic - Error when create new consumer %+v", err))

		return
	}

	listTopic, _ := consumer.Topics()
	fmt.Printf(fmt.Sprintf("get topics available in kafka: %s", base.JSONDebugDataString(listTopic)))

	listTopicNotExisted := make([]string, 0)
	for _, topicName := range c.config.ProducerTopics {
		if _, found := checkExistTopic(listTopic, topicName); !found {
			listTopicNotExisted = append(listTopicNotExisted, topicName)
		}
	}

	if len(listTopicNotExisted) == 0 {
		fmt.Printf(fmt.Sprintf("ClientKafka::CreateTopic - All of %v was existed", c.config.ProducerTopics))

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
		fmt.Printf(fmt.Sprintf("ClientKafka::CreateTopic - CreateTopics Error: %v", err))

		return
	}

	t := response.TopicErrors
	for key, val := range t {
		if val.Err != sarama.ErrNoError {
			fmt.Printf(fmt.Sprintf("😡😡😡 ClientKafka::CreateTopic - Create topic key: %s - Error: %s at pod %s in host %s", key, val.Err.Error(), env.Environment.Server.Host, env.Environment.Server.Host))
		}
	}
}

// ProducerPushMessage func;
func (c *ClientKafka) ProducerPushMessage(topic string, messageObj MessageKafka) (partition int32, offset int64, err error) {
	// debug.PrintStack()
	if c == nil || c.producer == nil {
		return 0, 0, fmt.Errorf("ClientKafka::ProducerPushMessage - Not found any producer")
	}

	fmt.Printf(fmt.Sprintf("ClientKafka::ProducerPushMessage - Push to topic: %s object data: %s", topic, base.JSONDebugDataString(messageObj)))

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

	fmt.Printf(fmt.Sprintf("ClientKafka::ProducerPushMessage - Push to topic: %s, key: %s, object data: %s", topic, key, base.JSONDebugDataString(messageObj)))

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(base.JSONDebugDataString(messageObj)),
	}

	return c.producer.SendMessage(msg)
}
