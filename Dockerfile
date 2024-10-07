FROM golang:1.23 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

RUN go mod init isrposandbox.com/project

ENV CGO_ENABLED=0
RUN go get -v ./...

RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime
COPY --from=build /go/src/swagger ./
EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]