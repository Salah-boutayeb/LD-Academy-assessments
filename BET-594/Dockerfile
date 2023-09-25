FROM golang:1.20.2-alpine
WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . ./

RUN go build -o .

EXPOSE 8080

CMD [ "/app" ]

