FROM golang:1.22.5

WORKDIR /app

COPY . .

RUN rm -rf supernode
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o bin

CMD ["./bin"]