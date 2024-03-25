package main

import (
	"context"
	"fmt"
	"log"

	"github.com/lukirs95/monika-gosdk/pkg/mocks"
	"github.com/lukirs95/monika-gosdk/pkg/provider"
)

func main() {
	gatewayEndpoint := "http://127.0.0.1:8080"
	mockProvider := provider.NewMockProvider()

	service := provider.NewService(gatewayEndpoint, mockProvider, log.Default())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	devices := mockProvider.GetDevices()
	updateChan := make(chan interface{})
	for _, device := range devices {
		mockDevice := mocks.NewMockDevice(device, updateChan)
		go mockDevice.Connect(ctx)
	}

	err := service.Listen(ctx, 8090, updateChan)
	fmt.Println(err)
}
