package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/nsqio/go-nsq"
)

func produce(addrs []string) {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second
	producers := []*nsq.Producer{}
	for _, addr := range addrs {
		producer, err := nsq.NewProducer(addr, cfg)
		if err != nil {
			panic(err)
		}
		producers = append(producers, producer)
	}

	for {
		now := time.Now().Format(time.RFC3339)
		index := rand.Intn(len(producers))
		log.Printf("publish via %d: %s", index, now)
		err := producers[index].Publish("events", []byte(now))
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
	}
}

func main() {
	// produce message via nsqd_1
	number, _ := strconv.Atoi(os.Args[1])
	nsqd := []string{}
	for i := 0; i < number; i++ {
		nsqd = append(nsqd, fmt.Sprintf("nsqd-cluster_nsqd_%d:4150", i+1))
	}

	go produce(nsqd)

	select {}

	// // produce message via nsqd_2
	// go produce("nsqd-cluster_nsqd_2:4150")

	// // produce message via nsqd_2
	// go produce("nsqd-cluster_nsqd_3:4150")
}
