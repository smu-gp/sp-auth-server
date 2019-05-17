package main

import (
	"github.com/go-redis/redis"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"net"

	envConfig "github.com/smu-gp/sp-sync-server/config/env"
	connectionDeliveryGrpc "github.com/smu-gp/sp-sync-server/connection/delivery/grpc"
	connectionRepository "github.com/smu-gp/sp-sync-server/connection/repository"
	connectionUsecase "github.com/smu-gp/sp-sync-server/connection/usecase"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	config := envConfig.NewViperConfig()
	serverAddr := config.GetString(`server.addr`)
	redisAddr := config.GetString(`database.redis.addr`)
	redisDb := config.GetInt(`database.redis.db`)

	sugar.Infof("SERVER_ADDR=%s", serverAddr)
	sugar.Infof("REDIS_ADDR=%s", redisAddr)
	sugar.Infof("REDIS_DB=%d", redisDb)

	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		sugar.Fatalf("Failed to listen: %v", err)
	}
	sugar.Infof("Start listening %s", serverAddr)

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   redisDb,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		sugar.Fatalf("Failed to connect redis: %v", err)
	}
	defer redisClient.Close()
	sugar.Infof("Connect redis: %s", redisAddr)

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_zap.StreamServerInterceptor(logger))),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(grpc_zap.UnaryServerInterceptor(logger))))

	connRepository := connectionRepository.NewRedisConnectionRepository(redisClient)
	connUsecase := connectionUsecase.NewConnectionUsecase(connRepository)
	connectionDeliveryGrpc.NewConnectionGrpcServer(grpcServer, connUsecase)

	reflection.Register(grpcServer)

	sugar.Info("Start serve grpc server")
	err = grpcServer.Serve(lis)
	if err != nil {
		sugar.Fatalf("Failed to serve: %v", err)
	}
}
