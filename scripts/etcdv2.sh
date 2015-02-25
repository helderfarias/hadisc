#!/bin/bash

docker run -d --name etcd coreos/etcd:v2.0.0 \
--listen-peer-urls 'http://0.0.0.0:2380,http://0.0.0.0:7001' \
--listen-client-urls 'http://0.0.0.0:2379,http://0.0.0.0:4001' \
--initial-advertise-peer-urls 'http://0.0.0.0:2380,http://0.0.0.0:7001' \
--initial-cluster 'default=http://0.0.0.0:2380,default=http://0.0.0.0:7001' \
--initial-cluster-state 'new' \
--initial-cluster-token 'etcd-cluster-1' \
--advertise-client-urls 'http://0.0.0.0:2379,http://0.0.0.0:4001'
