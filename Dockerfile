FROM golang:1.13.3-stretch
MAINTAINER cbping  452775680@qq.com
WORKDIR /go/src/app
ENV PATH /go/src/app:$PATH
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o gitk8s  .
RUN chmod +x gitk8s
CMD ["./gitk8s"]