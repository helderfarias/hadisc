FROM golang

RUN apt-get update && apt-get install -y libssl1.0.0 --no-install-recommends && rm -rf /var/lib/apt/lists/*

# Install haproxy
ENV HAPROXY_MAJOR 1.5
ENV HAPROXY_VERSION 1.5.11
ENV HAPROXY_MD5 5500a79d0d2b238d4a1e9749bd0c2cb2
RUN buildDeps='curl gcc libc6-dev libpcre3-dev libssl-dev make' \
	&& set -x \
	&& apt-get update && apt-get install -y $buildDeps --no-install-recommends && rm -rf /var/lib/apt/lists/* \
	&& curl -SL "http://www.haproxy.org/download/${HAPROXY_MAJOR}/src/haproxy-${HAPROXY_VERSION}.tar.gz" -o haproxy.tar.gz \
	&& echo "${HAPROXY_MD5}  haproxy.tar.gz" | md5sum -c \
	&& mkdir -p /usr/src/haproxy \
	&& tar -xzf haproxy.tar.gz -C /usr/src/haproxy --strip-components=1 \
	&& rm haproxy.tar.gz \
	&& make -C /usr/src/haproxy \
		TARGET=linux2628 \
		USE_PCRE=1 PCREDIR= \
		USE_OPENSSL=1 \
		USE_ZLIB=1 \
		all \
		install-bin \
	&& mkdir -p /usr/local/etc/haproxy \
	&& cp -R /usr/src/haproxy/examples/errorfiles /usr/local/etc/haproxy/errors \
	&& rm -rf /usr/src/haproxy \
	&& apt-get purge -y --auto-remove $buildDeps

# Install haproxy
RUN apt-get update \
    && apt-get install -y ca-certificates nginx \
    && rm /etc/nginx/sites-available/* \
    && rm /etc/nginx/sites-enabled/*


RUN apt-get update && apt-get -y install supervisor

RUN go get github.com/tools/godep
RUN go install github.com/tools/godep

RUN go get github.com/helderfarias/hadisc
RUN cd src/github.com/helderfarias/hadisc && godep go build
RUN cd src/github.com/helderfarias/hadisc \
    && cp hadisc /usr/bin/hadisc \
    && chmod +x /usr/bin/hadisc

RUN cd $GOPATH/src/github.com/helderfarias/hadisc \
    && mkdir -p /etc/haproxy/template \
    && cp templates/haproxy.tpl /etc/haproxy/template

RUN cd $GOPATH/src/github.com/helderfarias/hadisc \
    && mkdir -p /etc/nginx/template \
    && cp templates/nginx.tpl /etc/nginx/template

RUN cd $GOPATH/src/github.com/helderfarias/hadisc \
    && cp templates/supervisord.conf /etc/supervisor/conf.d/supervisord.conf 

EXPOSE 8080
EXPOSE 8090

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]

