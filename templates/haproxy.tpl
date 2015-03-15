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
  bind *:8080
  {{range $host := .}} 
  acl host_{{$host.Domain}} path_reg ^{{$host.Endpoint}}?
  {{end}}
  {{range $host := .}} 
  use_backend {{$host.Domain}} if host_{{$host.Domain}}
  {{end}}

{{range $domain := .}} 
backend {{$domain.Domain}}
  mode http
  balance roundrobin
  {{range $domain.Backends}}
  server {{.Container}} {{.Server}} check inter 2s rise 3 fall 2
  {{end}}
{{end}}
