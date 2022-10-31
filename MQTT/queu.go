package MQTT

import "github.com/Kephas73/lib-kephas/base"

const (
	KMqttDefaultQoS     = 0
	KMqttDefaultQuiesce = 250
)

func (manager *MQTTClient) Publish(topic string, message MessageQueue) error {
	token := manager.Client.Publish(topic, KMqttDefaultQoS, false, base.JSONDebugDataString(message))
	token.Wait()
	err := token.Error()
	if err != nil {
		return err
	}
	return nil
}
