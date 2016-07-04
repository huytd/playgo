FROM golang:alpine

ADD . $GOPATH/src/github.com/huytd/playgo
RUN mkdir /app
RUN ln -s $GOPATH/src/github.com/huytd/playgo/www /app/www
RUN chdir $GOPATH/src/github.com/huytd/playgo && go build -o /app/playgo .

EXPOSE 3000

CMD ["/app/playgo", "-mode=web"]
