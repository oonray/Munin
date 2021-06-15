FROM golang:1.16
RUN mkdir -p /ddns

WORKDIR /ddns
COPY . .
RUN go get github.com/sirupsen/logrus

CMD ["ddns"]
