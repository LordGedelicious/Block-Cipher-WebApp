FROM golang:1.19-alpine AS backend
WORKDIR /go/src/block-cipher-webapp/backend
RUN chmod +X ./
COPY backend/ ./
RUN go mod download
RUN go build -o backend .

CMD ["./backend"]