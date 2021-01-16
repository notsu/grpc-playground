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
