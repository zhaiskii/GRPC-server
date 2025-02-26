package main

import(
	"log"
	"net"
	"google.golang.org/grpc"
	pb "yalyceum_hw_grpc/pkg/api/test/api" 
	"yalyceum_hw_grpc/internal/repo"
	"yalyceum_hw_grpc/internal/service"
	"google.golang.org/grpc/reflection"
)

func main(){
	lis, err := net.Listen("tcp", "localhost:50051")
	//also smth with this error
	if err!=nil {
		//smth smth
	}
	
	orderRepo := repo.NewOrderRepo()
	orderService := service.NewOrderServiceServer(orderRepo)

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, orderService)

	reflection.Register(server	)
	log.Println("gRPC server running on port 50051")

	err2 := server.Serve(lis)
	if err2!=nil{
		//smth
	} 
}
