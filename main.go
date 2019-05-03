package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net"

	envConfig "github.com/smu-gp/sp-sync-server/config/env"
	connectionDeliveryGrpc "github.com/smu-gp/sp-sync-server/connection/delivery/grpc"
	connectionRepository "github.com/smu-gp/sp-sync-server/connection/repository"
	connectionUsecase "github.com/smu-gp/sp-sync-server/connection/usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := envConfig.NewViperConfig()
	serverAddr := config.GetString(`server.addr`)
	redisAddr := config.GetString(`database.redis.addr`)
	redisDb := config.GetInt(`database.redis.db`)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   redisDb,
	})
	defer redisClient.Close()

	connRepository := connectionRepository.NewRedisConnectionRepository(redisClient)
	connUsecase := connectionUsecase.NewConnectionUsecase(connRepository)

	grpcServer := grpc.NewServer()
	connectionDeliveryGrpc.NewConnectionGrpcServer(grpcServer, connUsecase)
	fmt.Println("Running server at", serverAddr)

	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
