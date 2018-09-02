FROM golang:latest
ENV gopa /app/go
RUN mkdir -p /app
ADD . /app 
WORKDIR /app/go/src/server 

RUN export GOPATH=${gopa} && echo ${GOPATH} && go get server 
RUN export GOPATH=${gopa} && go build -o server -ldflags "-X main.configFile=conf_docker.json" main.go 
EXPOSE 8000 
EXPOSE 443
CMD ["./server"]
