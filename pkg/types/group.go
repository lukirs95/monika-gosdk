package types

type Group struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GroupMember struct {
	Id         int64      `json:"id,omitempty"`
	ModuleId   ModuleId   `json:"moduleId,omitempty"`
	ModuleName string     `json:"modulename,omitempty"`
	ModuleType ModuleType `json:"moduletype,omitempty"`
	DeviceId   DeviceId   `json:"deviceId,omitempty"`
	DeviceName string     `json:"devicename,omitempty"`
	DeviceType DeviceType `json:"devicetype,omitempty"`
	Group      int64      `json:"group,omitempty"`
}
