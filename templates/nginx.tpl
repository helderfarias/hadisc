{{range $domain := .}} 
upstream {{$domain.Domain}} {  
   {{range $domain.Backends}}
   server {{.Server}};
   {{end}}
}
{{end}}

server {

    listen       8080;
    server_name  localhost;

    {{range $host := .}} 
    location ^~ {{$host.Endpoint}} {
       proxy_pass http://{{$host.Domain}};
    }
    {{end}}
}
