package provider

import (
	"fmt"

	"github.com/lukirs95/monika-gosdk/pkg/mocks"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type mockProvider struct {
	devices []types.Device
}

func NewMockProvider() Provider {
	devices := make([]types.Device, 0)
	for i := 0; i < 10; i++ {
		devices = append(devices, mocks.GetMockDevice(fmt.Sprint(i), fmt.Sprintf("Mock Device %d", i)))
	}

	return &mockProvider{
		devices: devices,
	}
}

func (provider *mockProvider) GetDeviceType() types.DeviceType {
	return types.DeviceType__GENERIC_DUMMY
}

func (provider *mockProvider) GetDevices() []types.Device {
	return provider.devices
}

func (m *mockProvider) GetDevice(deviceId types.DeviceId) types.Device {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device
		}
	}
	return nil
}
func (m *mockProvider) RunDeviceControl(deviceId types.DeviceId, cmd types.DeviceControl) error {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.FireAction(cmd)
		}
	}
	return fmt.Errorf("device not found")
}

func (m *mockProvider) GetModuleTypes(deviceId types.DeviceId) []types.ModuleType {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModuleTypes()
		}
	}
	return nil
}

func (m *mockProvider) GetModules(deviceId types.DeviceId) []types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModules()
		}
	}
	return nil
}

func (m *mockProvider) GetModulesByModuleType(deviceId types.DeviceId, moduleType types.ModuleType) []types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModulesByType(moduleType)
		}
	}
	return nil
}

func (m *mockProvider) GetModule(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModule(moduleId)
		}
	}
	return nil
}

func (m *mockProvider) RunModuleControl(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, cmd types.ModuleControl) error {
	var module types.Module
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			module = device.GetModule(moduleId)
		}
	}

	if module == nil {
		return fmt.Errorf("module not found")
	}

	return module.FireAction(cmd)
}

func (m *mockProvider) GetIOletTypes(deviceId types.DeviceId, moduleId types.ModuleId) []types.IOletType {
	var module types.Module
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			module = device.GetModule(moduleId)
		}
	}

	if module == nil {
		return nil
	}

	return module.GetIOletTypes()
}

func (m *mockProvider) GetIOlets(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) []types.IOlet {
	var module types.Module
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			module = device.GetModule(moduleId)
		}
	}

	if module == nil {
		return nil
	}

	return module.GetIOlets()
}

func (m *mockProvider) GetIOletsByIOletType(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType) []types.IOlet {
	var module types.Module
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			module = device.GetModule(moduleId)
		}
	}

	if module == nil {
		return nil
	}

	return module.GetIOletsByType(ioletType)
}

func (m *mockProvider) GetIOlet(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType, ioletId types.IOletId) types.IOlet {
	var module types.Module
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			module = device.GetModule(moduleId)
		}
	}

	if module == nil {
		return nil
	}

	return module.GetIOlet(ioletId)
}

func (m *mockProvider) RunIOletCommand(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType, ioletId types.IOletId, cmd types.IOletControl) error {
	var module types.Module
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			module = device.GetModule(moduleId)
		}
	}

	if module == nil {
		return nil
	}

	if iolet := module.GetIOlet(ioletId); iolet != nil {
		return iolet.FireAction(cmd)
	}

	return fmt.Errorf("iolet not found")
}

func (m *mockProvider) GetUpdates() []types.DeviceUpdate {
	updatedDevices := make([]types.DeviceUpdate, 0)
	for _, device := range m.devices {
		updatedDevice := device.Updated()
		if updatedDevice != nil {
			updatedDevices = append(updatedDevices, *updatedDevice)
		}
	}

	return updatedDevices
}
