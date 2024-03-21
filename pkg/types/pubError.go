package types

type PubError struct {
	ErrorId    int64            `json:"errorId,omitempty"`
	DeviceType DeviceType       `json:"deviceType"`
	DeviceName string           `json:"deviceName"`
	ModuleType ModuleType       `json:"moduleType,omitempty"`
	ModuleName string           `json:"moduleName,omitempty"`
	Severity   PubErrorSeverity `json:"severity"`
	Message    string           `json:"message"`
}

type PubErrorSeverity string

const (
	PubErrorSeverity_LOWEST  PubErrorSeverity = "LOWEST"
	PubErrorSeverity_MID     PubErrorSeverity = "MID"
	PubErrorSeverity_HIGHEST PubErrorSeverity = "HIGHEST"
)

type PubErrorResponse struct {
	ErrorId int64 `json:"errorId"`
}
