package main

import (
	"log"
	"time"

	"github.com/youzan/go-nsq"
)

func produce(addr string, seeds []string) {
	cfg := nsq.NewConfig()
	cfg.LookupdSeeds = seeds

	producer, err := nsq.NewProducer(addr, cfg)
	if err != nil {
		panic(err)
	}

	for {
		now := time.Now().Format(time.RFC3339)
		log.Printf("%s publish: %s", addr, now)
		err = producer.Publish("events", []byte(now))
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}

func main() {
	seeds := []string{
		"nsqd-cluster_nsqlookupd_1:4161",
		"nsqd-cluster_nsqlookupd_2:4161",
		"nsqd-cluster_nsqlookupd_3:4161",
	}

	// produce message via nsqd_1
	go produce("nsqd-cluster_nsqd_1:4150", seeds)

	// produce message via nsqd_2
	go produce("nsqd-cluster_nsqd_2:4150", seeds)

	go produce("nsqd-cluster_nsqd_3:4150", seeds)

	// go produce(seeds[0], seeds)
	// go produce(seeds[1], seeds)
	// go produce(seeds[2], seeds)
	select {}
}
