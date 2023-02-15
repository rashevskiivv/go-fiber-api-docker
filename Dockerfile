FROM golang
WORKDIR /go/src/go-fiber-api-docker
COPY . .
RUN go mod download
COPY . /go/src/go-fiber-api-docker
RUN go build -o bin/go-fiber-api-docker ./cmd/main.go
EXPOSE 8080 8080
CMD ["./bin/go-fiber-api-docker"]