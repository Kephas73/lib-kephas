package mqtt_client

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type MQTTClient struct {
	Client mqtt.Client
}

var mqInstance *MQTTClient

// default value env key is "MQTT";
// if configKeys was set, key env will be first value (not empty) of this;
func InstanceMQTTClientManager(configKeys ...string) *MQTTClient {
	if mqInstance != nil {
		return mqInstance
	}

	if config == nil {
		getMQTTConfigFromEnv(configKeys...)
	}

	if config == nil {
		err := fmt.Errorf("need config for mqtt client first")
		panic(err)
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", config.Host, config.Port))
	opts.SetUsername(config.Username)
	opts.SetPassword(config.Password)
	opts.SetClientID(fmt.Sprintf("%s_%d", config.ClientID, time.Now().UnixNano()))

	mqInstance = &MQTTClient{
		Client: mqtt.NewClient(opts),
	}

	token := mqInstance.Client.Connect()

	retry := 0
	for !token.WaitTimeout(10 * time.Second) {
		retry++
		if retry >= 3 {
			break
		}
	}

	if err := token.Error(); err != nil {
		panic(err)
	}

	return mqInstance
}

func GetMQTTClientInstance() *MQTTClient {
	if mqInstance == nil {
		return InstanceMQTTClientManager()
	}

	return mqInstance
}
