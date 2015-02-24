frontend http-in
  bind *:80
  {{range $host := .}} 
  acl host_{{$host.Domain}} path_beg ^/api/v2/{{$host.Domain}}
  {{end}}
  {{range $backend := .}} 
  use_backend {{$backend.Domain}}
  {{end}}

{{range $domain := .}} 
backend {{$domain.Domain}}
  mode http
  balance roundrobin
  {{range $domain.Backends}}
  server {{.Container}} {{.Server}} check inter 2s rise 3 fall 2
  {{end}}
{{end}}