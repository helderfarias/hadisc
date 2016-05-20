#!/bin/bash

docker run -d -p 2380:2380 -p 2379:2379 --name etcd quay.io/coreos/etcd:v2.0.0 -name infra0 \
 -initial-advertise-peer-urls http://0.0.0.0:2380 \
 -listen-client-urls http://0.0.0.0:2379 \
 -listen-peer-urls http://0.0.0.0:2380 \
 -advertise-client-urls http://0.0.0.0:2379 \
 -initial-cluster-token etcd-cluster-1 \
 -initial-cluster infra0=http://0.0.0.0:2380 \
 -initial-cluster-state new
