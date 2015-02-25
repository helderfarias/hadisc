#!/bin/bash

ETCD_HOST=10.89.4.165
CAD=localhost
ADMIN=localhost

curl -XDELETE http://$ETCD_HOST:4001/v2/keys/services?recursive=true&dir=true
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/cad/domain -d value=/api/v1/cad
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/cad/backend/container1 -d value=$CAD:3100
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/admin/domain -d value=/api/v1/admin
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/admin/backend/container2 -d value=$ADMIN:3100