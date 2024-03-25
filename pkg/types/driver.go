package types

type Driver struct {
	DeviceType DeviceType `json:"deviceType"`
	Port       int        `json:"port"`
	Location   string     `json:"location"`
}
