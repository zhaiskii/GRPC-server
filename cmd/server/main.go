package main

import(
	"log"
	"net"
	"net/http"
	"os"
	"fmt"
	"yalyceum_hw_grpc/internal/config"
	"os/signal"
	"syscall"
	"time"
	"google.golang.org/grpc"
	"context"
	pb "yalyceum_hw_grpc/pkg/api/test/api" 
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"yalyceum_hw_grpc/internal/repo"
	"yalyceum_hw_grpc/internal/service"
	"google.golang.org/grpc/reflection"
	"github.com/ilyakaznacheev/cleanenv"
)

func main(){
	sigChan := make(chan os.Signal, 1)//WTF IS HERE
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)//WTF IS HERE


	cfg := config.New()
	err_nafig := cleanenv.ReadConfig("./config.yaml", &cfg)
	_ = err_nafig
	//fmt.Println(cfg.Port_grpc)
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
	go func() {
		log.Println("gRPC server running on port 50051")
		err2 := server.Serve(lis)
		if err2!=nil{
			//smth
		} 
	}()
	mux:=runtime.NewServeMux()
	httpServer := &http.Server{
		Addr: ":8082",
		Handler: mux,
	}
	go func() {
		
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err = pb.RegisterOrderServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
		if err!=nil{
			fmt.Errorf("main.go registering http server %s", err)
		}
		log.Println("running server on 8082")
		if err:=httpServer.ListenAndServe(); err!=nil{
			fmt.Errorf("http server listen and serve main.go %s", err)
		}
	}()

	<-sigChan
	log.Println("Server Stopped")
	server.GracefulStop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = httpServer.Shutdown(ctx)
	if err!=nil{
		log.Fatalf("forced to shutdown %v", err)
	}
	log.Println("successfully stopped")
}
