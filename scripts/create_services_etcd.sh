#!/bin/bash

ETCD_HOST=192.168.99.100
CONTAINER_IP_1=192.168.99.100
CONTAINER_IP_2=192.168.99.100

curl -XDELETE http://$ETCD_HOST:4001/v2/keys/services?recursive=true&dir=true

curl -XPOST   http://$ETCD_HOST:4001/v2/keys/services/app01/domain -d value=/api
curl -XPOST   http://$ETCD_HOST:4001/v2/keys/services/app01/backend/one -d value=192.168.99.100:49159
curl -XPOST   http://$ETCD_HOST:4001/v2/keys/services/app01/backend/two -d value=192.168.99.100:49160

curl -XPOST   http://$ETCD_HOST:4001/v2/keys/services/app02/domain -d value=/icone
curl -XPOST   http://$ETCD_HOST:4001/v2/keys/services/app02/backend/three -d value=192.168.99.100:49163
curl -XPOST   http://$ETCD_HOST:4001/v2/keys/services/app02/backend/four -d value=192.168.99.100:49164
