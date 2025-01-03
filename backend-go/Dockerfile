# Build the Go application
FROM golang:1.23-alpine AS build  
WORKDIR /app  
COPY go.mod go.sum ./  
RUN go mod download  
COPY . .  
RUN go build -o /app/main .  

# Final runtime image
FROM alpine:latest  
WORKDIR /app  
COPY --from=build /app/main /app/main  
EXPOSE 8081  
CMD ["/main"]  
