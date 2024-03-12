package types

type Module struct {
	Type    ModuleType   `json:"type"`
	Name    string       `json:"name"`
	Status  ModuleStatus `json:"status"`
	Inlets  []IOlet      `json:"inlets,omitempty"`
	Outlets []IOlet      `json:"outlets,omitempty"`
}

type ModuleType string

const (
	ModuleType_IP    ModuleType = "IP"
	ModuleType_BB    ModuleType = "BB"
	ModuleType_POWER ModuleType = "POWER"
)

type ModuleStatus uint8

const (
	ModuleStatus_OK ModuleStatus = 0b1
)

func (s ModuleStatus) OK() bool {
	return s&ModuleStatus_OK != 0
}
