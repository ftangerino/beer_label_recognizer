FROM golang:1.20

WORKDIR /goapp

COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 5001

CMD ["./main"]
