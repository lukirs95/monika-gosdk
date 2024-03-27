package types

import (
	"fmt"
	"sync/atomic"
)

type IOlet interface {
	GetId() IOletId
	GetName() string
	GetType() IOletType
	SetName(newName string)
	GetStatus() IOletStatus
	SetStatus(newStatus IOletStatus)
	AddAction(control IOletControl, action IOletAction)
	FireAction(control IOletControl) error
	Updated() *IOletUpdate
}

func NewIOlet(id IOletId, ioletType IOletType, name string) IOlet {
	return &ioletImpl{
		Id:       id,
		Type:     ioletType,
		Name:     name,
		Status:   0,
		Controls: make([]IOletControl, 0),
		actions:  make(map[IOletControl]IOletAction),
		modified: atomic.Bool{},
	}
}

type ioletImpl struct {
	Id       IOletId        `json:"id"`
	Type     IOletType      `json:"type"`
	Name     string         `json:"name"`
	Status   IOletStatus    `json:"status"`
	Controls []IOletControl `json:"controls"`
	actions  map[IOletControl]IOletAction
	modified atomic.Bool
}

type IOletUpdate struct {
	Id     IOletId     `json:"id"`
	Type   IOletType   `json:"type"`
	Name   string      `json:"name"`
	Status IOletStatus `json:"status"`
}

func (iolet *ioletImpl) GetId() IOletId {
	return iolet.Id
}

func (iolet *ioletImpl) GetName() string {
	return iolet.Name
}

func (iolet *ioletImpl) GetType() IOletType {
	return iolet.Type
}

func (iolet *ioletImpl) SetName(newName string) {
	if iolet.Name != newName && newName != "" {
		iolet.Name = newName
		iolet.modified.Store(true)
	}
}

func (iolet *ioletImpl) GetStatus() IOletStatus {
	return iolet.Status
}
func (iolet *ioletImpl) SetStatus(newStatus IOletStatus) {
	if iolet.Status != newStatus {
		iolet.Status = newStatus
		iolet.modified.Store(true)
	}
}

func (iolet *ioletImpl) AddAction(newControl IOletControl, action IOletAction) {
	iolet.actions[newControl] = action

	for _, control := range iolet.Controls {
		if control == newControl {
			return
		}
	}
	iolet.Controls = append(iolet.Controls, newControl)
}

func (iolet *ioletImpl) FireAction(control IOletControl) error {
	if action, ok := iolet.actions[control]; ok {
		return action(iolet)
	}
	return fmt.Errorf("no such action defined")
}

func (iolet *ioletImpl) Updated() *IOletUpdate {
	if iolet.modified.Swap(false) {
		return &IOletUpdate{
			Id:     iolet.Id,
			Type:   iolet.Type,
			Name:   iolet.Name,
			Status: iolet.Status,
		}
	}
	return nil
}

type IOletId string

type IOletType string

const (
	IOletType_IPVIDEOIN  IOletType = "IP-VIDEO-IN"
	IOletType_IPVIDEOOUT IOletType = "IP-VIDEO-OUT"
	IOletType_IPAUDIOIN  IOletType = "IP-AUDIO-IN"
	IOletType_IPAUDIOOUT IOletType = "IP-AUDIO-OUT"
	IOletType_IPDATA     IOletType = "IP-DATA"
	IOletType_IPTIMING   IOletType = "IP-TIMING"
	IOletType_IPGPIO     IOletType = "IP-GPI"
	IOletType_BBVIDEOIN  IOletType = "BB-VIDEO-IN"
	IOletType_BBVIDEOOUT IOletType = "BB-VIDEO-OUT"
	IOletType_BBAUDIOIN  IOletType = "BB-AUDIO-IN"
	IOletType_BBAUDIOOUT IOletType = "BB-AUDIO-OUT"
	IOletType_BBTIMING   IOletType = "BB-TIMING"
	IOletType_BBGPIO     IOletType = "BB-GPI"
)

type IOletStatus uint8

const (
	IOletStatus_NOK       = 0b1 << 0
	IOletStatus_RUNNING   = 0b1 << 1
	IOletStatus_RECEIVING = 0b1 << 2
	IOletStatus_SENDING   = 0b1 << 3
	IOletStatus_HIGH      = 0b1 << 4
	IOletStatus_ENABLED   = 0b1 << 5
)

func (s IOletStatus) OK() bool {
	return s&IOletStatus_NOK == 0
}

func (s *IOletStatus) SetOK(ok bool) {
	setFlag(s, IOletStatus_NOK, !ok)
}

func (s IOletStatus) Running() bool {
	return s&IOletStatus_RUNNING > 0
}

func (s *IOletStatus) SetRunning(running bool) {
	setFlag(s, IOletStatus_RUNNING, running)
}

func (s IOletStatus) Receiving() bool {
	return s&IOletStatus_RECEIVING > 0
}

func (s *IOletStatus) SetReceiving(receiving bool) {
	setFlag(s, IOletStatus_RECEIVING, receiving)
}

func (s IOletStatus) Sending() bool {
	return s&IOletStatus_SENDING > 0
}

func (s *IOletStatus) SetSending(sending bool) {
	setFlag(s, IOletStatus_SENDING, sending)
}

func (s IOletStatus) High() bool {
	return s&IOletStatus_HIGH > 0
}

func (s *IOletStatus) SetHigh(high bool) {
	setFlag(s, IOletStatus_HIGH, high)
}

func (s IOletStatus) Enabled() bool {
	return s&IOletStatus_ENABLED > 0
}

func (s *IOletStatus) SetEnabled(enabled bool) {
	setFlag(s, IOletStatus_ENABLED, enabled)
}

func setFlag(reg *IOletStatus, flag IOletStatus, v bool) {
	if v {
		*reg |= flag
	} else {
		*reg &= ^flag
	}
}

type IOletControl string

const (
	IOletControl_START   IOletControl = "START"
	IOletControl_STOP    IOletControl = "STOP"
	IOletControl_RESTART IOletControl = "RESTART"
)

func (control IOletControl) Valid() error {
	valid := control == IOletControl_START || control == IOletControl_STOP || control == IOletControl_RESTART
	if valid {
		return nil
	}
	return fmt.Errorf("%s is not a valid control", control)
}

type IOletAction func(iolet IOlet) error
