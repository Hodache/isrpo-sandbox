FROM golang:latest AS build

WORKDIR /go/src

COPY go.mod ./
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0

RUN go mod tidy && go get -d -v ./...
RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime

COPY --from=build /go/src/swagger ./

EXPOSE 8080/tcp

ENTRYPOINT ["./swagger"]