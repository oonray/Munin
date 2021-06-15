FROM golang
RUN mkdir -p /ddns
RUN go get github.com/sirupsen/logrus
COPY main.go /ddns/main.go
WORKDIR /ddns
CMD ["bash","-c","go run main.go"]
