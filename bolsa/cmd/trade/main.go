package main

import (
	"encoding/json"
	"fmt"
	"sync"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/luigi20/bolsa_de_valores/bolsa/internal/infra/kafka"
	"github.com/luigi20/bolsa_de_valores/bolsa/internal/market/dto"
	"github.com/luigi20/bolsa_de_valores/bolsa/internal/market/entity"
	"github.com/luigi20/bolsa_de_valores/bolsa/internal/market/transformer"
)

func main() {
	ordersIn := make(chan *entity.Order)
	ordersOut := make(chan *entity.Order)
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	kafkaMsgChan := make(chan *ckafka.Message)
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "localhost:9094",
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	}
	producer := kafka.NewKafkaProducer(configMap)
	kafka := kafka.NewConsumer(configMap, []string{"input"})
	go kafka.Consume(kafkaMsgChan) // Tread 2

	//recebe do canal do kafka, joga no input, processa joga no output e depois publica no kafka
	book := entity.NewBook(ordersIn, ordersOut, wg)
	go book.Trade() // Tread 3

	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1)
			fmt.Println(string(msg.Value))
			tradeInput := dto.TradeInput{}
			err := json.Unmarshal(msg.Value, &tradeInput)
			if err != nil {
				panic(err)
			}
			order := transformer.TransformInput(tradeInput)
			ordersIn <- order
		}
	}()

	for res := range ordersOut {
		output := transformer.TransformOutput(res)
		outputJSON, err := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(outputJSON))
		if err != nil {
			fmt.Println(err)
		}
		producer.Publish(outputJSON, []byte("orders"), "output")
	}
}
