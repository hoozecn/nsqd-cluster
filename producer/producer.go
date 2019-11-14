package main

import (
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func produce(addr string) {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		panic(err)
	}

	for {
		now := time.Now().Format(time.RFC3339)
		log.Printf("publish: %s", now)
		err = producer.Publish("events", []byte(now))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	// produce message via nsqd_1
	go produce("nsqd-cluster_nsqd_1:4150")

	// produce message via nsqd_2
	go produce("nsqd-cluster_nsqd_2:4150")
	select {}
}
