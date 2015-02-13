global
  log 127.0.0.1   local0
  log 127.0.0.1   local1 notice
  maxconn 4096
  user root
  group root

defaults
  log    global
  mode   http
  option httplog
  option dontlognull
  option forwardfor
  option http-server-close
  option httpclose
  timeout connect 5000
  timeout client  50000
  timeout server  50000

frontend http-in
  bind *:80
  default_backend services

backend services
  mode http
  balance roundrobin
  {{range .}}
  server {{.Container}} {{.Server}} check inter 2s rise 3 fall 2
  {{end}}