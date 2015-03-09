# hadisc

### Service Discovery (HAProxy + NGINX + ETCD)

```
   docker run --name discovery -d --link etcd:etcd -e ETCD_HOST_DISCOVERY=http://etcd:2379 hadisc
```