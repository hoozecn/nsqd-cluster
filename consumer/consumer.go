package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

type handler struct {
	name string
}

func (h *handler) HandleMessage(message *nsq.Message) error {
	log.Printf("%s receive message %s", h.name, string(message.Body))
	return nil
}

func consume(addrs []string, name string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("events", name+"#ephemeral", cfg)
	if err != nil {
		panic(err)
	}

	// consumer.AddHandler(&handler{})
	consumer.AddConcurrentHandlers(&handler{name: name}, 1)

	err = consumer.ConnectToNSQLookupds(addrs)
	if err != nil {
		panic(err)
	}

	select {}
}

func main() {
	addrs := []string{
		"nsqd-cluster_nsqlookupd_1:4161",
		"nsqd-cluster_nsqlookupd_2:4161",
		"nsqd-cluster_nsqlookupd_3:4161",
	}
	go consume(addrs, "consumer_1")
	go consume(addrs, "consumer_2")
	select {}
}
