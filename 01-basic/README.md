# Generate/Regenerate gRPC code

## Ping Service

```
cd ping-service/proto

protoc --go_out=plugins=grpc:. greet.proto
```

## Pong Service

```
cd pong-service/proto

protoc --go_out=plugins=grpc:. greet.proto
```

## Getting Started

```
docker-compose up
```

## Table of Contents

1. Unary
2. Streaming on the server
3. Streaming on the client
4. Bi-direction streaming
