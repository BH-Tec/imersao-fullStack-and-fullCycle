package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/BH-Tec/imersao-fullStack-and-fullCycle/application/route"
	"github.com/BH-Tec/imersao-fullStack-and-fullCycle/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go"
)

// {"clientId":"1", "routeId":"1"}
// {"clientId":"2", "routeId":"2"}
// {"clientId":"3", "routeId":"3"}
func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()

	json.Unmarshal(msg.Value, &route)
	route.LoadPosition()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
