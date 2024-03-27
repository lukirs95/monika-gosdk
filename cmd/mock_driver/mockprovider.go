package main

import (
	"context"
	"fmt"

	"github.com/lukirs95/monika-gosdk/pkg/provider"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type MockProvider struct {
	deviceType types.DeviceType
	devices    []types.Device
	length     int
}

func NewMockProvider(deviceType types.DeviceType, length int) provider.DeviceProvider {

	return &MockProvider{
		deviceType: deviceType,
		devices:    make([]types.Device, 0),
		length:     length,
	}
}

func (provider *MockProvider) FetchDevices(ctx context.Context) error {
	for i := 0; i < provider.length; i++ {
		provider.devices = append(provider.devices, types.NewDevice(types.DeviceId(fmt.Sprint(i)), provider.deviceType, fmt.Sprintf("Mock Device %d", i)))
	}
	return nil
}

func (provider *MockProvider) GetDevices() []types.Device {
	return provider.devices
}

func (provider *MockProvider) GetDeviceType() types.DeviceType {
	return types.DeviceType__GENERIC_DUMMY
}
