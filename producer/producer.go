package main

import (
	"log"
	"time"

	"github.com/youzan/go-nsq"
)

func produce(addr string) {
	cfg := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		panic(err)
	}

	for {
		now := time.Now().Format(time.RFC3339)
		log.Printf("%s publish: %s", addr, now)
		err = producer.Publish("messages", []byte(now))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	// produce message via nsqd_1
	go produce("nsqd-cluster_nsqd_1:4150")

	// produce message via nsqd_2
	go produce("nsqd-cluster_nsqd_2:4150")

	go produce("nsqd-cluster_nsqd_3:4150")
	select {}
}
