FROM golang:1.18
# editer 
WORKDIR /go/src
COPY . . 
RUN go build cmd/main.go

EXPOSE 8080 
