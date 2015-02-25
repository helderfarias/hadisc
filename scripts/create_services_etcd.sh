#!/bin/bash

ETCD_HOST=$(docker inspect -f "{{.NetworkSettings.IPAddress}}" etcd)

curl -XDELETE http://$ETCD_HOST:4001/v2/keys/services?recursive=true&dir=true
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/cad/domain -d value=cad
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/admin/domain -d value=admin
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/cad/backend/one -d value=localhost:3000
curl -XPOST http://$ETCD_HOST:4001/v2/keys/services/admin/backend/two -d value=localhost:3001