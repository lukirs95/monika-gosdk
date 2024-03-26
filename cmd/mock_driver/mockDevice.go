package main

import (
	"context"
	"fmt"
	"time"

	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type MockDevice struct {
	device        types.Device
	triggerUpdate chan types.Device
}

func NewMockDevice(device types.Device, updateChan chan types.Device) *MockDevice {
	mock := &MockDevice{
		device:        device,
		triggerUpdate: updateChan,
	}

	device.AddAction(types.DeviceControl_BOOT, mock.deviceActionBoot)
	device.AddAction(types.DeviceControl_SHUTDOWN, mock.deviceActionShutDown)
	device.AddAction(types.DeviceControl_REBOOT, mock.deviceActionReboot)
	return mock
}

func (mock *MockDevice) Connect(ctx context.Context) {
	mockVideoModule := types.NewModule("1", types.ModuleType_AV, "Video Module")
	mockVideoModule.AddIOlet(types.NewIOlet("1", types.IOletType_IPVIDEOIN, "Video Input 1"))
	mockVideoModule.AddIOlet(types.NewIOlet("2", types.IOletType_IPVIDEOIN, "Video Input 2"))
	mockVideoModule.AddIOlet(types.NewIOlet("3", types.IOletType_IPVIDEOIN, "Video Input 3"))
	mockVideoModule.AddIOlet(types.NewIOlet("4", types.IOletType_IPVIDEOIN, "Video Input 4"))
	mockVideoModule.AddIOlet(types.NewIOlet("5", types.IOletType_IPVIDEOOUT, "Video Output 1"))
	mockVideoModule.AddIOlet(types.NewIOlet("6", types.IOletType_IPVIDEOOUT, "Video Output 2"))
	mockVideoModule.AddIOlet(types.NewIOlet("7", types.IOletType_IPVIDEOOUT, "Video Output 3"))
	mockVideoModule.AddIOlet(types.NewIOlet("8", types.IOletType_IPVIDEOOUT, "Video Output 4"))
	mock.device.AddModule(mockVideoModule)
	mockAudioModule := types.NewModule("2", types.ModuleType_AV, "Audio Module")
	mockAudioModule.AddIOlet(types.NewIOlet("1", types.IOletType_IPAUDIOIN, "Audio Input 1"))
	mockAudioModule.AddIOlet(types.NewIOlet("2", types.IOletType_IPAUDIOIN, "Audio Input 2"))
	mockAudioModule.AddIOlet(types.NewIOlet("3", types.IOletType_IPAUDIOIN, "Audio Input 3"))
	mockAudioModule.AddIOlet(types.NewIOlet("4", types.IOletType_IPAUDIOIN, "Audio Input 4"))
	mockAudioModule.AddIOlet(types.NewIOlet("5", types.IOletType_IPAUDIOOUT, "Audio Output 1"))
	mockAudioModule.AddIOlet(types.NewIOlet("6", types.IOletType_IPAUDIOOUT, "Audio Output 2"))
	mockAudioModule.AddIOlet(types.NewIOlet("7", types.IOletType_IPAUDIOOUT, "Audio Output 3"))
	mockAudioModule.AddIOlet(types.NewIOlet("8", types.IOletType_IPAUDIOOUT, "Audio Output 4"))
	mock.device.AddModule(mockAudioModule)

	for i := 0; i < 10; i++ {
		mockModule := types.NewModule(types.ModuleId(fmt.Sprint(i)), types.ModuleType_AV, fmt.Sprintf("Mock Module %d", i))

		for i := 0; i < 20; i++ {
			var mockIOlet types.IOlet
			if i%2 == 0 {
				mockIOlet = types.NewIOlet(types.IOletId(fmt.Sprint(i)), types.IOletType_IPVIDEOIN, fmt.Sprintf("Video In %d", i))
			} else {
				mockIOlet = types.NewIOlet(types.IOletId(fmt.Sprint(i)), types.IOletType_IPVIDEOOUT, fmt.Sprintf("Video Out %d", i))
			}

			mockModule.AddIOlet(mockIOlet)
		}
		mock.device.AddModule(mockModule)
	}

	for _, module := range mock.device.GetModules() {
		module.AddAction(types.ModuleControl_START, mock.moduleActionStart)
		module.AddAction(types.ModuleControl_STOP, mock.moduleActionStop)
		for _, iolet := range module.GetIOlets() {
			iolet.AddAction(types.IOletControl_START, mock.ioletActionStart)
			iolet.AddAction(types.IOletControl_STOP, mock.ioletActionStop)
		}
	}

	updateTicker := time.NewTicker(time.Second)
	defer updateTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case tick := <-updateTicker.C:
			mock.mockUpdate(tick)
		}
	}
}

func (mock *MockDevice) mockUpdate(currentTime time.Time) {
	mock.device.SetName(currentTime.Format("15:04:05"))
	mock.triggerUpdate <- mock.device
}

func (mock *MockDevice) deviceActionBoot(device types.Device) error {
	currentStatus := mock.device.GetStatus()
	currentStatus.SetONLINE(true)
	mock.device.SetStatus(currentStatus)
	mock.triggerUpdate <- mock.device
	return nil
}

func (mock *MockDevice) deviceActionShutDown(device types.Device) error {
	currentStatus := mock.device.GetStatus()
	currentStatus.SetONLINE(false)
	mock.device.SetStatus(currentStatus)
	mock.triggerUpdate <- mock.device
	return nil
}

func (mock *MockDevice) deviceActionReboot(device types.Device) error {
	currentStatus := mock.device.GetStatus()
	currentStatus.SetONLINE(false)
	mock.device.SetStatus(currentStatus)
	mock.triggerUpdate <- mock.device
	go func(mock *MockDevice) {
		time.Sleep(time.Second * 5)
		currentStatus := mock.device.GetStatus()
		currentStatus.SetONLINE(true)
		mock.device.SetStatus(currentStatus)
		mock.triggerUpdate <- mock.device
	}(mock)
	return nil
}

func (mock *MockDevice) moduleActionStart(module types.Module) error {
	currentStatus := module.GetStatus()
	currentStatus.SetOK(true)
	module.SetStatus(currentStatus)
	mock.triggerUpdate <- mock.device
	return nil
}

func (mock *MockDevice) moduleActionStop(module types.Module) error {
	currentStatus := module.GetStatus()
	currentStatus.SetOK(false)
	module.SetStatus(currentStatus)
	mock.triggerUpdate <- mock.device
	return nil
}

func (mock *MockDevice) ioletActionStart(iolet types.IOlet) error {
	currentStatus := iolet.GetStatus()
	currentStatus.SetRunning(true)
	currentStatus.SetReceiving(true)
	iolet.SetStatus(currentStatus)
	mock.triggerUpdate <- mock.device
	return nil
}

func (mock *MockDevice) ioletActionStop(iolet types.IOlet) error {
	currentStatus := iolet.GetStatus()
	currentStatus.SetRunning(false)
	currentStatus.SetReceiving(false)
	iolet.SetStatus(currentStatus)
	mock.triggerUpdate <- mock.device
	return nil
}
