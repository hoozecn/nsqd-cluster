## 有赞NSQ集群评估 （需要安装 docker和docker-compose）

3个ectd节点，3个nsqlookupd，3个nsqd，1个consumer，1个producer

1. start nsq cluster

`sh start_services.sh`

2. start producer

`sh start_producer.sh`

3. start consumer

`sh start_consumer.sh`

4. visit nsqadmin http://127.0.0.1:24171

