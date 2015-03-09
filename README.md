# hadisc

### Service Discovery (HAProxy + ETCD)
```
   docker run --privileged \
      --name discovery \
      -d --link etcd:etcd \
      -e ETCD_HOST_DISCOVERY=http://etcd:2379 hadisc
```
