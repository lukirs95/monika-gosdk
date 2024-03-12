package types

type IOlet struct {
	Type     IOletType      `json:"type"`
	Name     string         `json:"name"`
	Status   IOletStatus    `json:"status"`
	Controls []IOletControl `json:"controls,omitempty"`
}

type IOletType string

const (
	IOletType_IPVIDEO  IOletType = "IP-VIDEO"
	IOletType_IPAUDIO  IOletType = "IP-AUDIO"
	IOletType_IPTIMING IOletType = "IP-TIMING"
	IOletType_IPGPIO   IOletType = "IP-GPIO"
	IOletType_BBVIDEO  IOletType = "BB-VIDEO"
	IOletType_BBAUDIO  IOletType = "BB-AUDIO"
	IOletType_BBTIMING IOletType = "BB-TIMING"
	IOletType_BBGPIO   IOletType = "BB-GPIO"
)

type IOletStatus uint8

const (
	IOletStatus_RUNNING   = 0b1 << 0
	IOletStatus_RECEIVING = 0b1 << 1
	IOletStatus_SENDING   = 0b1 << 2
	IOletStatus_HIGH      = 0b1 << 3
)

func (s IOletStatus) Running() bool {
	return s&IOletStatus_RUNNING != 0
}

func (s IOletStatus) Receiving() bool {
	return s&IOletStatus_RECEIVING != 0
}

func (s IOletStatus) Sending() bool {
	return s&IOletStatus_SENDING != 0
}

func (s IOletStatus) High() bool {
	return s&IOletStatus_HIGH != 0
}

type IOletControl string

const (
	IOletControl_STOP    IOletControl = "STOP"
	IOletControl_START   IOletControl = "START"
	IOletControl_RESTART IOletControl = "RESTART"
)
