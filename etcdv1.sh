docker run -d -p 4001:4001 -p 5001:7001 coreos/etcd -addr=0.0.0.0:4001 -peer-addr=0.0.0.0:5001
