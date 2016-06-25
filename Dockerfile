FROM golang:alpine

ADD . $GOPATH/src/github.com/huytd/go-play
RUN mkdir /app
RUN ln -s $GOPATH/src/github.com/huytd/go-play/www /app/www
RUN chdir $GOPATH/src/github.com/huytd/go-play && go build -o /app/playgo .

EXPOSE 3000

CMD ["/app/playgo", "-mode=web"]
