package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/uuid"
	envConfig "github.com/smu-gp/sp-sync-server/config/env"
	connectionGrpc "github.com/smu-gp/sp-sync-server/protobuf/connection"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"os"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	config := envConfig.NewViperConfig()
	serverAddr := config.GetString(`server.addr`)

	conn, err := grpc.Dial(`127.0.0.1`+serverAddr, grpc.WithInsecure())
	if err != nil {
		sugar.Error(`Failed connect`, err)
	}
	defer conn.Close()

	var userId = uuid.New().String()
	sugar.Infof("Generated userId: %s", userId)

	client := connectionGrpc.NewConnectionServiceClient(conn)
	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userId})
	sugar.Info(connectionResponse)

	go func() {
		stream, _ := client.WaitAuth(context.Background())
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId: userId,
		})
		response, _ := stream.Recv()
		sugar.Info("Auth deviceName: ", response.AuthDevice.DeviceName, ", deviceType: ", response.AuthDevice.DeviceType)
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId:       userId,
			AuthDevice:   response.AuthDevice,
			AcceptDevice: true,
		})
	}()

	code := promptConnectionCode()
	authResponse, _ := client.Auth(context.Background(), &connectionGrpc.AuthRequest{ConnectionCode: code, DeviceInfo: &connectionGrpc.AuthDeviceInfo{
		DeviceName: "TestDevice",
		DeviceType: connectionGrpc.AuthDeviceInfo_DEVICE_TABLET,
	}})
	sugar.Info(authResponse.GetMessage(), " userId: ", authResponse.GetUserId())
}

func promptConnectionCode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter connection code: ")
	code, _ := reader.ReadString('\n')
	return code
}
