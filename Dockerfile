FROM alpine:edge
MAINTAINER Jim Alateras <jima@comware.com.au>

ENV GOPATH /golang
ENV FILESERVER_HOME /fileserver
ENV PATH $FILESERVER_HOME:$GOPATH/bin:$PATH

ADD . $GOPATH/src/github.com/jalateras/fileserver

RUN \
  addgroup -S fileserver && \
  adduser -S -s /bin/bash -G fileserver fileserver

RUN \
  apk add --update bash wget git mercurial bzr go make && \
  cd $GOPATH/src/github.com/jalateras/fileserver && \
  make bootstrap build && \
  mkdir -p $FILESERVER_HOME && \
  cp ./fileserver $FILESERVER_HOME/   && \
  cp -R ./public $FILESERVER_HOME/ && \
  chown -R fileserver:fileserver $FILESERVER_HOME && \
  apk del --purge wget git mercurial bzr go  && \
  rm -rf /var/cache/apk/* /tmp/* /var/tmp/* $GOPATH

# Expose the http api port
EXPOSE 8080

USER fileserver
WORKDIR /fileserver
CMD ["./fileserver"]



