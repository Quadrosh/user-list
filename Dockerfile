FROM golang:1.17
ENV GO111MODULE=on

RUN mkdir /userlist
ADD . /userlist
WORKDIR /userlist
RUN go mod tidy
RUN go build -o /userlist/build/api/bin  /userlist/cmd/api/main.go
ENTRYPOINT /userlist/build/api/bin -dbhost=db -dbport=5432 -dbname=userlist -dbuser=postgres -dbpass=password  -port=8000

EXPOSE 8000