#!/bin/sh
set -e

# start etcd
docker-compose up -d etcd1 etcd2 etcd3
# start nslookupd
docker-compose scale nsqlookupd=3
# start nsqd
docker-compose scale nsqd=3
# start nsqadmin
docker-compose scale nsqadmin=1

# create topic events
sleep 10
curl 'http://127.0.0.1:24171/api/topics' -H content-type:'application/json' --data-binary '{"topic":"events","partition_num":"1","replicator":"3","retention_days":"","syncdisk":"","orderedmulti":"false","extend":"false"}'
