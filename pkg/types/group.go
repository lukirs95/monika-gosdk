package types

type Group struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GroupMember struct {
	Id         int64      `json:"id,omitempty"`
	ModuleName string     `json:"modulename,omitempty"`
	ModuleType ModuleType `json:"moduletype,omitempty"`
	DeviceName string     `json:"devicename,omitempty"`
	DeviceType DeviceType `json:"devicetype,omitempty"`
	Group      int64      `json:"group,omitempty"`
}
