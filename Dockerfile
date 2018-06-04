FROM golang:1.8

RUN mkdir -p /go/src/github.com/nylo-andry/movies-service
WORKDIR /go/src/github.com/nylo-andry/movies-service

COPY . .

RUN go get -d -v ./
RUN go install -v ./

CMD ["movies-service"]