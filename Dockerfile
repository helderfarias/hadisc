FROM golang

RUN apt-get update && apt-get install -y libssl1.0.0 --no-install-recommends && rm -rf /var/lib/apt/lists/*

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

RUN apt-get update && apt-get -y install supervisor

RUN go get github.com/tools/godep
RUN go install github.com/tools/godep

RUN mkdir -p $GOPATH/src/github.com/helderfarias/hadisc
COPY discovery $GOPATH/src/github.com/helderfarias/hadisc/discovery
COPY drive $GOPATH/src/github.com/helderfarias/hadisc/drive
COPY helper $GOPATH/src/github.com/helderfarias/hadisc/helper
COPY main.go $GOPATH/src/github.com/helderfarias/hadisc/
RUN godep go install github.com/helderfarias/hadisc
RUN cp bin/hadisc /usr/bin/hadisc && chmod +x /usr/bin/hadisc

WORKDIR /etc/haproxy
COPY templates/haproxy.tpl /etc/haproxy/template/
COPY templates/supervisord.conf /etc/supervisor/conf.d/supervisord.conf 

EXPOSE 8080

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]

