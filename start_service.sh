#!/bin/sh
set -e

docker-compose up -d etcd1 etcd2 etcd3
docker run --network nsqd-cluster_default buildpack-deps:curl sh -c 'curl -X POST "http://nsqd-cluster_etcd1_1:2379/v2/keys/NSQMetaData/etcd-cluster/Topics"'
docker-compose scale nsqlookupd=3
docker-compose scale nsqd=3
docker-compose scale nsqadmin=1
