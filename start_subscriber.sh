docker run -it --network nsqd-cluster_default -v $PWD:/app -w /app golang:1.12.9 go run -mod=vendor ./subscriber/... $@