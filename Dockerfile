FROM golang:1.13.1

COPY . /k8s-job
WORKDIR /k8s-job
RUN go build -o main ./app

CMD ["/k8s-job/main", "-sleepTime=15"]