package driver

import (
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type Driver interface {
	DeviceDriver
	ModuleDriver
	IOletDriver
}

type DeviceDriver interface {
	// returns the deviceType the driver is responsible for
	GetDeviceType() types.DeviceType
	// returns all devices the driver handles without the modules.
	GetDevices() []types.Device
	// returns one device based on the deviceId
	GetDevice(deviceId types.DeviceId) types.Device
	// RunDeviceControl executes the given control command
	RunDeviceControl(deviceId types.DeviceId, cmd types.DeviceControl) error
	// returns the moduleTypes the driver has in the system
	GetModuleTypes(deviceId types.DeviceId) []types.ModuleType
}

type ModuleDriver interface {
	// returns all modules the driver handles without IOlet's
	GetModules(deviceId types.DeviceId) []types.Module
	// returns all modules the driver handles based on the moduleType.
	GetModulesByModuleType(deviceId types.DeviceId, moduleType types.ModuleType) []types.Module
	// returns one module based on the moduleId
	GetModule(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) types.Module
	// RunModuleControl executes the given control command
	RunModuleControl(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, cmd types.ModuleControl) error
}

type IOletDriver interface {
	// returns the IOletTypes the driver has in the system
	GetIOletTypes(deviceId types.DeviceId, moduleId types.ModuleId) []types.IOletType
	// returns all IOlets the given module has
	GetIOlets(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId) []types.IOlet
	// returns all IOlets the given module has based on ioletType
	GetIOletsByIOletType(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioletType types.IOletType) []types.IOlet
	// returns one IOlet based on the ioletId
	GetIOlet(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioLetType types.IOletType, ioLetId types.IOletId) types.IOlet
	// RunIOletCommand executes the given control command
	RunIOletCommand(deviceId types.DeviceId, moduleType types.ModuleType, moduleId types.ModuleId, ioLetType types.IOletType, ioLetId types.IOletId, cmd types.IOletControl) error
}
