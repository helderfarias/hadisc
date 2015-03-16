# Service Discovery

## Dockerfile

```
  docker build -t hadisc .
```

## Run

```
   docker run -d --name discovery -e ETCD_HOST=http://localhost:2379 hadisc
```
