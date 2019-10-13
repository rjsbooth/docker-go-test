# Step 1 - builder
FROM golang:1.13 as builder
COPY . /src/docker-go-test
WORKDIR /src/docker-go-test
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-go-test

# Step 2 - image
FROM alpine:latest
WORKDIR /bin/
COPY --from=builder /src/docker-go-test/docker-go-test .
EXPOSE 8000
CMD ["./docker-go-test"]