FROM golang:latest
WORKDIR /home/tianyuan/go_demo/practice/gin-blog
COPY . /home/tianyuan/go_demo/practice/gin-blog
RUN go build ./cmd/main.go
EXPOSE 8000
ENTRYPOINT ["./main"]