package builder

//go:generate protoc -I . -I vendor --gogo_out=grpc:. trigger.proto
