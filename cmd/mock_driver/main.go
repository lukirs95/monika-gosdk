package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lukirs95/monika-gosdk/pkg/driver"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

func main() {
	gatewayEndpoint := "http://127.0.0.1:8080"
	mockProvider := NewMockProvider(types.DeviceType__GENERIC_DUMMY)
	mockProvider.FetchDevices(context.Background())
	mockDriver, err := driver.NewDriver(mockProvider)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	mockService := driver.NewService(gatewayEndpoint, mockDriver, log.Default())

	mockService.AddErrorCheckIOlet(func(iolet *types.IOletUpdate) *types.Error {
		if !iolet.Status.Running() {
			return &types.Error{
				Severity: types.IOletStatus_HIGH,
				Message:  fmt.Sprintf("IOlet %s stopped!", iolet.Name),
			}
		}
		return nil
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	devices := mockDriver.GetDevices()
	updateChan := make(chan types.Device, 100)
	for _, device := range devices {
		mockDevice := NewMockDevice(device, updateChan)
		go mockDevice.Connect(ctx)
	}

	fmt.Print(mockService.Listen(ctx, 8090, updateChan))
}
