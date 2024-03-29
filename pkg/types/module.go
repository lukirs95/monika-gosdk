package types

import (
	"context"
	"fmt"
	"sync/atomic"
)

type Module interface {
	GetId() ModuleId
	GetType() ModuleType
	SetName(newName string)
	GetStatus() ModuleStatus
	SetStatus(newStatus ModuleStatus)
	AddAction(newControl ModuleControl, action ModuleAction)
	FireAction(ctx context.Context, control ModuleControl) error
	AddIOlet(newIOlet IOlet)
	GetIOletTypes() []IOletType
	GetIOlets() []IOlet
	GetIOletsByType(ioletType IOletType) []IOlet
	GetIOlet(ioletId IOletId) IOlet
	Updated() *ModuleUpdate
}

func NewModule(id ModuleId, moduleType ModuleType, name string) Module {
	return &moduleImpl{
		Id:         id,
		Type:       moduleType,
		Name:       name,
		Controls:   make([]ModuleControl, 0),
		actions:    make(map[ModuleControl]ModuleAction),
		IOletTypes: make([]IOletType, 0),
		IOlets:     make([]IOlet, 0),
		modified:   atomic.Bool{},
	}
}

type moduleImpl struct {
	Id         ModuleId        `json:"id"`
	Type       ModuleType      `json:"type"`
	Name       string          `json:"name"`
	Status     ModuleStatus    `json:"status"`
	Controls   []ModuleControl `json:"controls"`
	actions    map[ModuleControl]ModuleAction
	IOletTypes []IOletType `json:"ioletTypes"`
	IOlets     []IOlet     `json:"-"`
	modified   atomic.Bool
}

type ModuleUpdate struct {
	Id     ModuleId      `json:"id"`
	Type   ModuleType    `json:"type"`
	Name   string        `json:"name"`
	Status ModuleStatus  `json:"status"`
	IOlets []IOletUpdate `json:"iolets"`
}

func (module *moduleImpl) GetId() ModuleId {
	return module.Id
}

func (module *moduleImpl) GetType() ModuleType {
	return module.Type
}

func (module *moduleImpl) SetName(newName string) {
	if module.Name == newName && newName != "" {
		module.Name = newName
		module.modified.Store(true)
	}
}

func (module *moduleImpl) GetStatus() ModuleStatus {
	return module.Status
}

func (module *moduleImpl) SetStatus(newStatus ModuleStatus) {
	if module.Status != newStatus {
		module.Status = newStatus
		module.modified.Store(true)
	}
}

func (module *moduleImpl) AddAction(newControl ModuleControl, action ModuleAction) {
	module.actions[newControl] = action

	for _, control := range module.Controls {
		if control == newControl {
			return
		}
	}
	module.Controls = append(module.Controls, newControl)
}

func (module *moduleImpl) FireAction(ctx context.Context, control ModuleControl) error {
	if action, ok := module.actions[control]; ok {
		return action(ctx, module)
	}
	return fmt.Errorf("no such action defined")
}

func (module *moduleImpl) GetIOletTypes() []IOletType {
	return module.IOletTypes
}

func (module *moduleImpl) addIOletType(newIOletType IOletType) {
	for _, ioletType := range module.IOletTypes {
		if newIOletType == ioletType {
			return
		}
	}
	module.IOletTypes = append(module.IOletTypes, newIOletType)
}

func (module *moduleImpl) AddIOlet(newIOlet IOlet) {
	module.addIOletType(newIOlet.GetType())
	module.IOlets = append(module.IOlets, newIOlet)
}

func (module *moduleImpl) GetIOlets() []IOlet {
	return module.IOlets
}

func (module *moduleImpl) GetIOletsByType(ioletType IOletType) []IOlet {
	ioletsByType := make([]IOlet, 0)
	for _, iolet := range module.IOlets {
		if iolet.GetType() == ioletType {
			ioletsByType = append(ioletsByType, iolet)
		}
	}

	return ioletsByType
}

func (module *moduleImpl) GetIOlet(ioletId IOletId) IOlet {
	for _, iolet := range module.IOlets {
		if iolet.GetId() == ioletId {
			return iolet
		}
	}
	return nil
}

func (module *moduleImpl) Updated() *ModuleUpdate {
	updatedIOlets := make([]IOletUpdate, 0)
	for _, iolet := range module.IOlets {
		updated := iolet.Updated()
		if updated != nil {
			updatedIOlets = append(updatedIOlets, *updated)
		}
	}

	if module.modified.Swap(false) || len(updatedIOlets) > 0 {
		return &ModuleUpdate{
			Id:     module.Id,
			Type:   module.Type,
			Name:   module.Name,
			Status: module.Status,
			IOlets: updatedIOlets,
		}
	}
	return nil
}

func (module *moduleImpl) Modified() bool {
	return module.modified.Swap(false)
}

type ModuleId string

type ModuleType string

const (
	ModuleType_AV    ModuleType = "AV"
	ModuleType_GPIO  ModuleType = "GPIO"
	ModuleType_BB    ModuleType = "TIMING"
	ModuleType_POWER ModuleType = "POWER"
)

type ModuleStatus uint8

const (
	ModuleStatus_NOK ModuleStatus = 0b1
)

func (s ModuleStatus) OK() bool {
	return s&ModuleStatus_NOK == 0
}

func (s *ModuleStatus) SetOK(ok bool) {
	if ok {
		*s &= ^ModuleStatus_NOK
	} else {
		*s |= ModuleStatus_NOK
	}
}

type ModuleControl string

const (
	ModuleControl_START   ModuleControl = "START"
	ModuleControl_STOP    ModuleControl = "STOP"
	ModuleControl_RESTART ModuleControl = "RESTART"
)

func (control ModuleControl) Valid() error {
	valid := control == ModuleControl_START || control == ModuleControl_STOP || control == ModuleControl_RESTART
	if valid {
		return nil
	}
	return fmt.Errorf("%s is not a valid control", control)
}

type ModuleAction func(ctx context.Context, module Module) error
