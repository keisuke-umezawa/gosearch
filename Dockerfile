FROM golang:onbuild
MAINTAINER Masashi Shibata<contact@keisuke-umezawa.link>

ADD . $GOPATH/src/github.com/keisuke-umezawa/gosearch
WORKDIR $GOPATH/src/github.com/keisuke-umezawa/gosearch

RUN godep restore
RUN go install github.com/keisuke-umezawa/gosearch

# Run the outyet command by default when the container starts.
ENV GOSEARCH_ENV develop
ENTRYPOINT $GOPATH/bin/gosearch

EXPOSE 8080
