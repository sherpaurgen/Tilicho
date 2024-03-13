#pulling go image which has all dependencies, the image bigger than alpine
FROM golang:1.22.1 AS builder
# crete /app directory ,copy the source dirs,files to the /app  and build a compiled executable binary with file named "app" in same "/app" dir
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

#pull a minimal Docker image based on Alpine Linux which is just 5MB in size
FROM alpine:latest AS production
# copy he "app" binary from builder stage above
COPY --from=builder /app . 
CMD ["./app"]
# run the command
#docker build -t rest-api-v2 .