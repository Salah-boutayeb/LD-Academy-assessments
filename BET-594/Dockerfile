# FROM golang:1.20.2-alpine
# WORKDIR /app

# COPY go.mod .
# RUN go mod download

# COPY . ./

# RUN go build -o .

# EXPOSE 8080

# CMD [ "/app" ]

FROM golang:1.20.2 as builder
WORKDIR /app
COPY . .

RUN go mod init api
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gin-gonic/gin
RUN go get github.com/joho/godotenv
RUN go get golang.org/x/crypto/bcrypt
RUN go get gorm.io/gorm
RUN go get gorm.io/driver/postgres

RUN go build -o api

FROM alpine:3.14

WORKDIR /app
COPY --from=builder /app/ .
RUN chmod +x /app
EXPOSE 8080
CMD ["/app"]