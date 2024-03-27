package types

type PubError struct {
	ErrorId    int64            `json:"errorId"`
	DeviceId   DeviceId         `json:"deviceId"`
	DeviceType DeviceType       `json:"deviceType"`
	DeviceName string           `json:"deviceName"`
	ModuleId   ModuleId         `json:"moduleId,omitempty"`
	ModuleType ModuleType       `json:"moduleType,omitempty"`
	ModuleName string           `json:"moduleName,omitempty"`
	IOletId    IOletId          `json:"ioletId,omitempty"`
	IOletType  IOletType        `json:"ioletType,omitempty"`
	IOletName  string           `json:"ioletName,omitempty"`
	Severity   PubErrorSeverity `json:"severity"`
	Message    string           `json:"message"`
}

type PubErrorSeverity int8

const (
	PubErrorSeverity_LOWEST  PubErrorSeverity = 0
	PubErrorSeverity_MID     PubErrorSeverity = 5
	PubErrorSeverity_HIGHEST PubErrorSeverity = 10
)

type PubErrorResponse struct {
	ErrorId int64 `json:"errorId"`
}

type Error struct {
	ErrorId  int64            `json:"errorId"`
	Severity PubErrorSeverity `json:"severity"`
	Message  string           `json:"message"`
}

type ErrorCheckerDevice func(device *DeviceUpdate) *Error
type ErrorCheckerModule func(device *ModuleUpdate) *Error
type ErrorCheckerIOlet func(device *IOletUpdate) *Error
