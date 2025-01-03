# Build the Go application
FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build com otimização para evitar binário inflado
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# Final runtime image - Usa scratch (imagem vazia)
FROM scratch
WORKDIR /app
COPY --from=build /app/main /main

EXPOSE 8081
CMD ["/main"]
