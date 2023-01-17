package po

import (
	"fmt"
	"strconv"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Sing struct {
	Topic   string
	Payload float64
}

type Sub struct {
	Broker string
	Port   int
	Topics []*Sing
	QOS    int
}

func (s *Sub) Activate() {

	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Println(string(msg.Payload()), " from ", msg.Topic())
		pay := string(msg.Payload())
		for _, v := range s.Topics {
			if v.Topic == string(msg.Topic()) {
				v.Payload, _ = strconv.ParseFloat(pay, 64)
			}
		}
	}

	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("Connected")
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("Connect lost: %v", err)
	}

	var broker = s.Broker
	var port = s.Port
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("hydro")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for _, v := range s.Topics {
		Subscribe(client, v.Topic)
		fmt.Println("SUBED: " + v.Topic)
	}

	for true {

	}
	// time.Sleep(5 * time.Second)
}

func Subscribe(client mqtt.Client, top string) {
	topic := top
	token := client.Subscribe(topic, 0, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}
