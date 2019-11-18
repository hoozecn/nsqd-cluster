package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

type handler struct {
	name     string
	handling int32
	messages chan string
}

func (h *handler) HandleMessage(message *nsq.Message) error {
	select {
	case h.messages <- string(message.Body):
		// do nothing
	case <-time.After(100 * time.Millisecond):
		log.Printf("%s failed to consume message in 100ms", h.name)
	}
	return nil
}

func consume(addrs []string, name string) {
	cfg := nsq.NewConfig()
	// set a higher maxInFlight value to acceept message concurrently
	// !!! IMPORTANT
	cfg.MaxInFlight = 3
	cfg.LookupdPollInterval = time.Second
	consumer, err := nsq.NewConsumer("events", name+"#ephemeral", cfg)
	// consumer, err := nsq.NewConsumer("events", "events", cfg)
	if err != nil {
		panic(err)
	}

	// consumer.AddHandler(&handler{})
	messages := make(chan string, 2)
	go func() {
		for msg := range messages {
			log.Printf("%s receive message %s", name, msg)
			time.Sleep(400 * time.Millisecond)
			log.Printf("%s finish message %s", name, msg)
		}
	}()
	consumer.AddConcurrentHandlers(&handler{name: name, messages: messages}, 3)

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

	number, _ := strconv.Atoi(os.Args[1])

	for i := 0; i < number; i++ {
		go consume(addrs, "consumer_"+strconv.Itoa(i))
	}

	select {}
}
