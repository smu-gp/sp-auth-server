package main

import (
	"bufio"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	envConfig "github.com/smu-gp/sp-sync-server/config/env"
	connectionGrpc "github.com/smu-gp/sp-sync-server/protobuf/build"
	"google.golang.org/grpc"
	"os"
)

func main() {
	config := envConfig.NewViperConfig()
	serverAddr := config.GetString(`server.addr`)

	conn, err := grpc.Dial(`127.0.0.1`+serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Error(`Failed connect`, err)
	}
	defer conn.Close()

	client := connectionGrpc.NewConnectionServiceClient(conn)
	userIdResponse, err := client.RequestUserId(context.Background(), &connectionGrpc.Empty{})
	if err != nil {
		log.Fatal(`Failed call request user id`)
	}
	log.Info(userIdResponse)

	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userIdResponse.UserId})
	log.Info(connectionResponse)

	go func() {
		stream, _ := client.WaitAuth(context.Background())
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId: userIdResponse.UserId,
		})
		response, _ := stream.Recv()
		log.Info("Auth deviceName: ", response.AuthDevice.DeviceName, ", deviceType: ", response.AuthDevice.DeviceType)
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId:       userIdResponse.UserId,
			AuthDevice:   response.AuthDevice,
			AcceptDevice: true,
		})
	}()

	code := promptConnectionCode()
	authResponse, _ := client.Auth(context.Background(), &connectionGrpc.AuthRequest{ConnectionCode: code, DeviceInfo: &connectionGrpc.AuthDeviceInfo{
		DeviceName: "TestDevice",
		DeviceType: connectionGrpc.AuthDeviceInfo_DEVICE_TABLET,
	}})
	log.Info(authResponse.GetMessage(), " userId: ", authResponse.GetUserId())
}

func promptConnectionCode() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter connection code: ")
	code, _ := reader.ReadString('\n')
	return code
}
