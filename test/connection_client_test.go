package main

import (
	"context"
	"github.com/google/uuid"
	connectionGrpc "github.com/smu-gp/sp-sync-server/protobuf/connection"
	"google.golang.org/grpc"
	"testing"
)

func InitConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(`127.0.0.1:8001`, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func TestConnection(t *testing.T) {
	conn, err := InitConn()
	defer conn.Close()
	if err != nil {
		t.Error("Not connected")
	}

	var userId = uuid.New().String()
	client := connectionGrpc.NewConnectionServiceClient(conn)
	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userId})
	if err != nil {
		t.Error("Error occurred")
	}
	if len(connectionResponse.ConnectionCode) == 0 {
		t.Error("Not generated connection code")
	}
}

func TestAuth_Success(t *testing.T) {
	conn, err := InitConn()
	defer conn.Close()
	if err != nil {
		t.Error("Not connected")
	}

	var userId = uuid.New().String()
	client := connectionGrpc.NewConnectionServiceClient(conn)
	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userId})
	if err != nil {
		t.Error("Error occurred")
	}
	if len(connectionResponse.ConnectionCode) == 0 {
		t.Error("Not generated connection code")
	}
	go func() {
		stream, _ := client.WaitAuth(context.Background())
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId: userId,
		})
		response, _ := stream.Recv()
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId:       userId,
			AuthDevice:   response.AuthDevice,
			AcceptDevice: true,
		})
	}()

	code := connectionResponse.ConnectionCode
	authResponse, _ := client.Auth(context.Background(), &connectionGrpc.AuthRequest{ConnectionCode: code, DeviceInfo: &connectionGrpc.AuthDeviceInfo{
		DeviceName: "TestDevice",
		DeviceType: connectionGrpc.AuthDeviceInfo_DEVICE_TABLET,
	}})

	if authResponse.Message != connectionGrpc.AuthResponse_MESSAGE_SUCCESS {
		t.Error("Failed auth")
	}
	if len(authResponse.UserId) == 0 {
		t.Error("Not exist userId")
	}
}

func TestAuth_Auth_Failed(t *testing.T) {
	conn, err := InitConn()
	defer conn.Close()
	if err != nil {
		t.Error("Not connected")
	}

	var userId = uuid.New().String()
	client := connectionGrpc.NewConnectionServiceClient(conn)
	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userId})
	if err != nil {
		t.Error("Error occurred")
	}
	if len(connectionResponse.ConnectionCode) == 0 {
		t.Error("Not generated connection code")
	}

	code := "0"
	authResponse, _ := client.Auth(context.Background(), &connectionGrpc.AuthRequest{ConnectionCode: code, DeviceInfo: &connectionGrpc.AuthDeviceInfo{
		DeviceName: "TestDevice",
		DeviceType: connectionGrpc.AuthDeviceInfo_DEVICE_TABLET,
	}})

	if authResponse.Message != connectionGrpc.AuthResponse_MESSAGE_FAILED {
		t.Error("Success auth")
	}
	if authResponse.FailedReason != connectionGrpc.AuthResponse_AUTH_FAILED {
		t.Error("FailedReason is not equal", authResponse.FailedReason)
	}
}

func TestAuth_Reject_Host(t *testing.T) {
	conn, err := InitConn()
	defer conn.Close()
	if err != nil {
		t.Error("Not connected")
	}

	var userId = uuid.New().String()
	client := connectionGrpc.NewConnectionServiceClient(conn)
	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userId})
	if err != nil {
		t.Error("Error occurred")
	}
	if len(connectionResponse.ConnectionCode) == 0 {
		t.Error("Not generated connection code")
	}
	go func() {
		stream, _ := client.WaitAuth(context.Background())
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId: userId,
		})
		response, _ := stream.Recv()
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId:       userId,
			AuthDevice:   response.AuthDevice,
			AcceptDevice: false,
		})
	}()

	code := connectionResponse.ConnectionCode
	authResponse, _ := client.Auth(context.Background(), &connectionGrpc.AuthRequest{ConnectionCode: code, DeviceInfo: &connectionGrpc.AuthDeviceInfo{
		DeviceName: "TestDevice",
		DeviceType: connectionGrpc.AuthDeviceInfo_DEVICE_TABLET,
	}})

	if authResponse.Message != connectionGrpc.AuthResponse_MESSAGE_FAILED {
		t.Error("Success auth")
	}
	if authResponse.FailedReason != connectionGrpc.AuthResponse_REJECT_HOST {
		t.Error("FailedReason is not equal", authResponse.FailedReason)
	}
}

func TestAuth_No_Host_Waited(t *testing.T) {
	conn, err := InitConn()
	defer conn.Close()
	if err != nil {
		t.Error("Not connected")
	}

	var userId = uuid.New().String()
	client := connectionGrpc.NewConnectionServiceClient(conn)
	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userId})
	if err != nil {
		t.Error("Error occurred")
	}
	if len(connectionResponse.ConnectionCode) == 0 {
		t.Error("Not generated connection code")
	}

	code := connectionResponse.ConnectionCode
	authResponse, _ := client.Auth(context.Background(), &connectionGrpc.AuthRequest{ConnectionCode: code, DeviceInfo: &connectionGrpc.AuthDeviceInfo{
		DeviceName: "TestDevice",
		DeviceType: connectionGrpc.AuthDeviceInfo_DEVICE_TABLET,
	}})

	if authResponse.Message != connectionGrpc.AuthResponse_MESSAGE_FAILED {
		t.Error("Success auth")
	}
	if authResponse.FailedReason != connectionGrpc.AuthResponse_NO_HOST_WAITED {
		t.Error("FailedReason is not equal", authResponse.FailedReason)
	}
}

func TestAuth_Response_Timeout(t *testing.T) {
	conn, err := InitConn()
	defer conn.Close()
	if err != nil {
		t.Error("Not connected")
	}

	var userId = uuid.New().String()
	client := connectionGrpc.NewConnectionServiceClient(conn)
	connectionResponse, err := client.Connection(context.Background(), &connectionGrpc.ConnectionRequest{UserId: userId})
	if err != nil {
		t.Error("Error occurred")
	}
	if len(connectionResponse.ConnectionCode) == 0 {
		t.Error("Not generated connection code")
	}

	code := connectionResponse.ConnectionCode
	go func() {
		stream, _ := client.WaitAuth(context.Background())
		_ = stream.Send(&connectionGrpc.WaitAuthRequest{
			UserId: userId,
		})
	}()

	authResponse, _ := client.Auth(context.Background(), &connectionGrpc.AuthRequest{ConnectionCode: code, DeviceInfo: &connectionGrpc.AuthDeviceInfo{
		DeviceName: "TestDevice",
		DeviceType: connectionGrpc.AuthDeviceInfo_DEVICE_TABLET,
	}})

	if authResponse.Message != connectionGrpc.AuthResponse_MESSAGE_FAILED {
		t.Error("Success auth")
	}
	if authResponse.FailedReason != connectionGrpc.AuthResponse_RESPONSE_TIMEOUT {
		t.Error("FailedReason is not equal", authResponse.FailedReason)
	}
}
