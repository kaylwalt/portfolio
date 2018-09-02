FROM golang:latest
ENV gopa /app/go
RUN mkdir -p /app
ADD . /app 
WORKDIR /app/go/src/server 

RUN export GOPATH=${gopa} && echo ${GOPATH} && go get server 
RUN export GOPATH=${gopa} && go build -o server main.go 
EXPOSE 8888 
CMD ["./server"]
