package types

import (
	"context"
	"encoding/json"
	"fmt"
	"sync/atomic"
)

type Device interface {
	SetId(newId DeviceId)
	GetId() DeviceId
	SetName(string)
	GetName() string
	GetStatus() DeviceStatus
	SetStatus(newStatus DeviceStatus)
	SetControlIP(controlIP string)
	GetControlIP() string
	SetControlPort(controlPort int)
	GetControlPort() int
	AddAction(control DeviceControl, callback DeviceAction)
	FireAction(ctx context.Context, control DeviceControl) error
	GetModuleTypes() []ModuleType
	AddModule(module Module)
	GetModules() []Module
	GetModulesByType(moduleType ModuleType) []Module
	GetModule(moduleId ModuleId) Module
	Updated() *DeviceUpdate
}

func NewDevice(id DeviceId, deviceType DeviceType, name string) Device {
	return &deviceImpl{
		Id:          id,
		Type:        deviceType,
		Name:        name,
		Status:      0,
		Controls:    make([]DeviceControl, 0),
		actions:     make(map[DeviceControl]DeviceAction),
		ModuleTypes: make([]ModuleType, 0),
		Modules:     make([]Module, 0),
		modified:    atomic.Bool{},
	}
}

func DevicesFromJSON(decoder *json.Decoder) ([]Device, error) {
	var rawDevices []deviceImpl
	if err := decoder.Decode(&rawDevices); err != nil {
		return nil, err
	}
	devices := make([]Device, 0)
	for index := range rawDevices {
		devices = append(devices, &rawDevices[index])
	}
	return devices, nil
}

func DeviceFromJSON(decoder *json.Decoder) (Device, error) {
	device := new(deviceImpl)
	if err := decoder.Decode(device); err != nil {
		return nil, err
	}
	return device, nil
}

type deviceImpl struct {
	Id          DeviceId        `json:"deviceId"`
	Type        DeviceType      `json:"type"`
	Name        string          `json:"name"`
	Status      DeviceStatus    `json:"status"`
	ControlIP   string          `json:"controlIP,omitempty"`
	ControlPort int             `json:"controlPort,omitempty"`
	Controls    []DeviceControl `json:"controls"`
	actions     map[DeviceControl]DeviceAction
	ModuleTypes []ModuleType `json:"moduleTypes"`
	Modules     []Module     `json:"-"`
	modified    atomic.Bool
}

type DeviceUpdate struct {
	Id      DeviceId       `json:"deviceId"`
	Type    DeviceType     `json:"type"`
	Name    string         `json:"name"`
	Status  DeviceStatus   `json:"status"`
	Modules []ModuleUpdate `json:"modules"`
}

func (device *deviceImpl) SetId(deviceId DeviceId) {
	device.Id = deviceId
}

func (device *deviceImpl) GetId() DeviceId {
	return device.Id
}

func (device *deviceImpl) SetName(newName string) {
	if device.Name != newName && newName != "" {
		device.Name = newName
		device.modified.Store(true)
	}
}

func (device *deviceImpl) GetName() string {
	return device.Name
}

func (device *deviceImpl) GetStatus() DeviceStatus {
	return device.Status
}

func (device *deviceImpl) SetStatus(newStatus DeviceStatus) {
	if device.Status != newStatus {
		device.Status = newStatus
		device.modified.Store(true)
	}
}

func (device *deviceImpl) SetControlIP(controlIP string) {
	device.ControlIP = controlIP
}

func (device *deviceImpl) GetControlIP() string {
	return device.ControlIP
}

func (device *deviceImpl) SetControlPort(controlPort int) {
	device.ControlPort = controlPort
}

func (device *deviceImpl) GetControlPort() int {
	return device.ControlPort
}

func (device *deviceImpl) AddAction(newControl DeviceControl, action DeviceAction) {
	device.actions[newControl] = action

	for _, control := range device.Controls {
		if control == newControl {
			return
		}
	}
	device.Controls = append(device.Controls, newControl)
}

func (device *deviceImpl) FireAction(ctx context.Context, control DeviceControl) error {
	if action, ok := device.actions[control]; ok {
		return action(ctx, device)
	}
	return fmt.Errorf("no such action defined")
}

func (device *deviceImpl) addModuleType(newModuleType ModuleType) {
	for _, moduleType := range device.ModuleTypes {
		if newModuleType == moduleType {
			return
		}
	}
	device.ModuleTypes = append(device.ModuleTypes, newModuleType)
}

func (device *deviceImpl) GetModuleTypes() []ModuleType {
	return device.ModuleTypes
}

func (device *deviceImpl) AddModule(module Module) {
	device.addModuleType(module.GetType())
	device.Modules = append(device.Modules, module)
}

func (device *deviceImpl) GetModules() []Module {
	return device.Modules
}

func (device *deviceImpl) GetModulesByType(moduleType ModuleType) []Module {
	modulesByType := make([]Module, 0)
	for _, module := range device.Modules {
		if module.GetType() == moduleType {
			modulesByType = append(modulesByType, module)
		}
	}

	return modulesByType
}

func (device *deviceImpl) GetModule(moduleId ModuleId) Module {
	for _, module := range device.Modules {
		if module.GetId() == moduleId {
			return module
		}
	}
	return nil
}

func (device *deviceImpl) Updated() *DeviceUpdate {
	updatedModules := make([]ModuleUpdate, 0)
	for _, module := range device.Modules {
		moduleUpdate := module.Updated()
		if moduleUpdate != nil {
			updatedModules = append(updatedModules, *moduleUpdate)
		}
	}

	if device.modified.Swap(false) || len(updatedModules) > 0 {
		return &DeviceUpdate{
			Id:      device.Id,
			Type:    device.Type,
			Name:    device.Name,
			Status:  device.Status,
			Modules: updatedModules,
		}
	}
	return nil
}

type DeviceId string

type DeviceType string

const (
	DeviceType__GENERIC_DUMMY  DeviceType = "GENERIC_DUMMY"
	DeviceType_GENERIC_USV     DeviceType = "GENERIC_USV"
	DeviceType_XLINK_XLINK     DeviceType = "XLINK_XLINK"
	DeviceType_RIEDEL_FUSION   DeviceType = "RIEDEL_FUSION"
	DeviceType_RIEDEL_MUON     DeviceType = "RIEDEL_MUON"
	DeviceType_RIEDEL_BOLERO   DeviceType = "RIEDEL_BOLERO"
	DeviceType_RIEDEL_NSA02    DeviceType = "RIEDEL_NSA02"
	DeviceType_DIRECTOUT_RAVIO DeviceType = "DIRECTOUT_RAVIO"
)

type DeviceStatus uint8

const (
	DeviceStatus_ONLINE DeviceStatus = 0b1
)

func (s DeviceStatus) ONLINE() bool {
	return s&DeviceStatus_ONLINE != 0
}

func (s *DeviceStatus) SetONLINE(online bool) {
	if online {
		*s |= DeviceStatus_ONLINE
	} else {
		*s &= ^DeviceStatus_ONLINE
	}
}

type DeviceControl string

const (
	DeviceControl_BOOT     DeviceControl = "BOOT"
	DeviceControl_REBOOT   DeviceControl = "REBOOT"
	DeviceControl_SHUTDOWN DeviceControl = "SHUTDOWN"
)

func (control DeviceControl) Valid() error {
	valid := control == DeviceControl_BOOT || control == DeviceControl_REBOOT || control == DeviceControl_SHUTDOWN
	if valid {
		return nil
	}
	return fmt.Errorf("%s is not a valid control", control)
}

type DeviceAction func(ctx context.Context, device Device) error
