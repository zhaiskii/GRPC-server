package interceptor

import (
    "context"
    "log"
    "google.golang.org/grpc"
)

func UnaryLoggingInterceptor(
    ctx context.Context, 
    req interface{}, 
    info *grpc.UnaryServerInfo, 
    handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("request %v", req)
	resp, err := handler(ctx, req)
	log.Printf("response %v", resp)
	return resp, err
}
