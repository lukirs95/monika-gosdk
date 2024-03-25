package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/lukirs95/monika-gosdk/pkg/externalprovider"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type providerImpl struct {
	devices          []types.Device
	gatewayEndpoint  string
	externalProvider externalprovider.DeviceProvider
}

func NewProvider(gatewayEndpoint string, externalProvider externalprovider.DeviceProvider) (Provider, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	devices, err := externalProvider.GetDevices(ctx)
	if err != nil {
		return nil, err
	}

	return &providerImpl{
		devices:          devices,
		externalProvider: externalProvider,
	}, nil
}

func (provider *providerImpl) GetDeviceType() types.DeviceType {
	return provider.externalProvider.GetDeviceType()
}

func (provider *providerImpl) GetDevices() []types.Device {
	return provider.devices
}

func (m *providerImpl) GetDevice(deviceId types.DeviceId) types.Device {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device
		}
	}
	return nil
}
func (m *providerImpl) RunDeviceControl(deviceId types.DeviceId, cmd types.DeviceControl) error {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.FireAction(cmd)
		}
	}
	return fmt.Errorf("device not found")
}

func (m *providerImpl) GetModuleTypes(deviceId types.DeviceId) []types.ModuleType {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModuleTypes()
		}
	}
	return nil
}

func (m *providerImpl) GetModules(deviceId types.DeviceId) []types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModules()
		}
	}
	return nil
}

func (m *providerImpl) GetModulesByModuleType(deviceId types.DeviceId, moduleType types.ModuleType) []types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModulesByType(moduleType)
		}
	}
	return nil
}

func (m *providerImpl) GetModule(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) types.Module {
	for _, device := range m.devices {
		if device.GetId() == deviceId {
			return device.GetModule(moduleId)
		}
	}
	return nil
}

func (m *providerImpl) RunModuleControl(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, cmd types.ModuleControl) error {
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

func (m *providerImpl) GetIOletTypes(deviceId types.DeviceId, moduleId types.ModuleId) []types.IOletType {
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

func (m *providerImpl) GetIOlets(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) []types.IOlet {
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

func (m *providerImpl) GetIOletsByIOletType(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType) []types.IOlet {
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

func (m *providerImpl) GetIOlet(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType, ioletId types.IOletId) types.IOlet {
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

func (m *providerImpl) RunIOletCommand(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType, ioletId types.IOletId, cmd types.IOletControl) error {
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

func (m *providerImpl) SendUpdates() {
	updatedDevices := make([]types.DeviceUpdate, 0)
	for _, device := range m.devices {
		updatedDevice := device.Updated()
		if updatedDevice != nil {
			updatedDevices = append(updatedDevices, *updatedDevice)
		}
	}

	if len(updatedDevices) > 0 {
		body, err := json.Marshal(updatedDevices)
		if err != nil {
			log.Print(err)
			return
		}
		reader := bytes.NewReader(body)
		http.Post(fmt.Sprintf("%s/api/notify/update", m.gatewayEndpoint), "application/json", reader)
	}
}

func (m *providerImpl) GetUpdates() []types.DeviceUpdate {
	updatedDevices := make([]types.DeviceUpdate, 0)
	for _, device := range m.devices {
		updatedDevice := device.Updated()
		if updatedDevice != nil {
			updatedDevices = append(updatedDevices, *updatedDevice)
		}
	}

	return updatedDevices
}
