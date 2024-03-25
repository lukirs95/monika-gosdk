package netbox

import (
	"fmt"
	"strings"

	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type Device struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	PrimaryIP PrimaryIP `json:"primary_ip"`
}

type PrimaryIP struct {
	Family  int    `json:"family"`  // 4 || 6
	Address string `json:"address"` // XXX.XXX.XXX.XXX/XX
}

func (d Device) GetId() types.DeviceId {
	return types.DeviceId(fmt.Sprint(d.Id))
}

func (d Device) GetName() string {
	return d.Name
}

func (d Device) GetIP() string {
	address := strings.Split(d.PrimaryIP.Address, "/")[0]
	return address
}
