package types

type Device struct {
	Type        DeviceType   `json:"type"`
	Name        string       `json:"name"`
	Status      DeviceStatus `json:"status"`
	ControlIP   string       `json:"controlIP,omitempty"`
	ControlPort int          `json:"controlPort,omitempty"`
	Modules     []Module     `json:"modules,omitempty"`
}

type DeviceType string

const (
	DeviceType_GENERIC_USV       DeviceType = "GENERIC_USV"
	DeviceType_XLINK_XLINK       DeviceType = "XLINK_XLINK"
	DeviceType_RIEDEL_FUSION     DeviceType = "RIEDEL_FUSION"
	DeviceType_RIEDEL_MUON       DeviceType = "RIEDEL_MUON"
	DeviceType_RIEDEL_BOLERO     DeviceType = "RIEDEL_BOLERO"
	DeviceType_RIEDEL_NSA02      DeviceType = "RIEDEL_NSA02"
	DeviceType_DIRECTOUT_PRODIGY DeviceType = "DIRECTOUT_PRODIGY"
)

type DeviceStatus uint8

const (
	DeviceStatus_OK DeviceStatus = 0b1
)

func (s DeviceStatus) OK() bool {
	return s&DeviceStatus_OK != 0
}
