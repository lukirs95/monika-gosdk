package types

type PubError struct {
	DeviceType DeviceType       `json:"deviceType"`
	DeviceName string           `json:"deviceName"`
	ModuleType ModuleType       `json:"moduleType"`
	ModuleName string           `json:"moduleName"`
	Severity   PubErrorSeverity `json:"severity"`
	Message    string           `json:"message"`
}

type PubErrorSeverity string

const (
	PubErrorSeverity_LOWEST  PubErrorSeverity = "LOWEST"
	PubErrorSeverity_MID     PubErrorSeverity = "MID"
	PubErrorSeverity_HIGHEST PubErrorSeverity = "HIGHEST"
)
