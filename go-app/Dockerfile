# Используйте последнюю версию Go
FROM golang:latest AS build

# Установите рабочую директорию
WORKDIR /go/src

# Копируйте исходный код
COPY go.mod ./
COPY go ./go
COPY main.go .

# Установите переменную окружения для отключения cgo
ENV CGO_ENABLED=0

# Установите зависимости
RUN go mod tidy && go get -d -v ./...

# Соберите бинарный файл
RUN go build -a -installsuffix cgo -o swagger .

# Используйте scratch для создания минимального образа
FROM scratch AS runtime

# Копируйте собранный бинарный файл в образ
COPY --from=build /go/src/swagger ./

# Экспонируйте порт
EXPOSE 8080/tcp

# Укажите точку входа
ENTRYPOINT ["./swagger"]