package provider

import (
	"context"

	"github.com/lukirs95/monika_gosdk/pkg/provider/netbox"
	"github.com/lukirs95/monika_gosdk/pkg/types"
)

type DeviceProvider interface {
	GetDevices(context.Context) ([]types.Device, error)
}

func NewDeviceProviderNetbox(server string, apiKey string, deviceTypeID int) DeviceProvider {
	return netbox.NewNetbox(server, apiKey, deviceTypeID)
}
