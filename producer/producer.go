package main

import (
	"log"
	"time"

	"github.com/youzan/go-nsq"
)

func produce(addresses []string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	topics := []string{"events"}
	producer, err := nsq.NewTopicProducerMgr(topics, cfg)
	if err != nil {
		panic(err)
	}

	producer.AddLookupdNodes(addresses)

	for {
		now := time.Now().Format(time.RFC3339)
		log.Printf("publish: %s", now)
		err = producer.Publish("events", []byte(now))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	lookupdAddresses := []string{
		"nsqd-cluster_nsqlookupd_1:4161",
		"nsqd-cluster_nsqlookupd_2:4161",
		"nsqd-cluster_nsqlookupd_3:4161",
	}

	go produce(lookupdAddresses)
	select {}
}
