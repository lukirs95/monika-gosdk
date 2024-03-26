package driver

import (
	"fmt"

	"github.com/lukirs95/monika-gosdk/pkg/provider"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type driverImpl struct {
	devices  []types.Device
	provider provider.DeviceProvider
}

func NewDriver(provider provider.DeviceProvider) (Driver, error) {
	devices := provider.GetDevices()

	return &driverImpl{
		devices:  devices,
		provider: provider,
	}, nil
}

func (driver *driverImpl) GetDeviceType() types.DeviceType {
	return driver.provider.GetDeviceType()
}

func (driver *driverImpl) GetDevices() []types.Device {
	return driver.devices
}

func (m *driverImpl) GetDevice(deviceId types.DeviceId) types.Device {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device
		}
	}
	return nil
}
func (m *driverImpl) RunDeviceControl(deviceId types.DeviceId, cmd types.DeviceControl) error {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.FireAction(cmd)
		}
	}
	return fmt.Errorf("device not found")
}

func (m *driverImpl) GetModuleTypes(deviceId types.DeviceId) []types.ModuleType {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModuleTypes()
		}
	}
	return nil
}

func (m *driverImpl) GetModules(deviceId types.DeviceId) []types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModules()
		}
	}
	return nil
}

func (m *driverImpl) GetModulesByModuleType(deviceId types.DeviceId, moduleType types.ModuleType) []types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModulesByType(moduleType)
		}
	}
	return nil
}

func (m *driverImpl) GetModule(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModule(moduleId)
		}
	}
	return nil
}

func (m *driverImpl) RunModuleControl(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, cmd types.ModuleControl) error {
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

func (m *driverImpl) GetIOletTypes(deviceId types.DeviceId, moduleId types.ModuleId) []types.IOletType {
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

func (m *driverImpl) GetIOlets(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) []types.IOlet {
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

func (m *driverImpl) GetIOletsByIOletType(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType) []types.IOlet {
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

func (m *driverImpl) GetIOlet(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType, ioletId types.IOletId) types.IOlet {
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

func (m *driverImpl) RunIOletCommand(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType, ioletId types.IOletId, cmd types.IOletControl) error {
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
