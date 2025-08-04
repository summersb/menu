FROM golang:tip-bookworm

WORKDIR /app

COPY go.mod ./
RUN go mod tidy
RUN go mod download

COPY . .
RUN go build -o recipe-app .

EXPOSE 8080

CMD ["./recipe-app"]
