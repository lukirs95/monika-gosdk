package externalprovider

import (
	"context"

	"github.com/lukirs95/monika-gosdk/pkg/externalprovider/netbox"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type DeviceProvider interface {
	GetDevices(context.Context) ([]types.Device, error)
	GetDeviceType() types.DeviceType
}

func NewDeviceProviderNetbox(server string, apiKey string, deviceType types.DeviceType, deviceTypeID int) DeviceProvider {
	return netbox.NewNetbox(server, apiKey, deviceType, deviceTypeID)
}
