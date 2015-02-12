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
  contimeout 5000
  clitimeout 50000
  srvtimeout 50000

frontend http-in
  bind *:80
  {{range $k, $v := .}}
  use_backend {{$v.Domain}}
  {{end}}  

{{range $k, $v := .}}
backend {{$v.Domain}}
  balance leastconn
  {{range $backend := $v.Backends}}
  server {{$backend.Container}} {{$backend.Server}} check inter 2s rise 3 fall 2
  {{end}}
{{end}}