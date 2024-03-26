package provider

import (
	"context"

	"github.com/lukirs95/monika-gosdk/pkg/provider/netbox"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type DeviceProvider interface {
	FetchDevices(context.Context) error
	// GetDevices is called by the driver. A second call MUST return the same references!
	GetDevices() []types.Device
	// GetDeviceType returns the deviceType the whole driver will be responsible for.
	GetDeviceType() types.DeviceType
}

// NewDeviceProviderNetbox returns a DeviceProvider which fetches devices from a netbox instance.
func NewDeviceProviderNetbox(server string, apiKey string, deviceType types.DeviceType, deviceTypeID int) DeviceProvider {
	return netbox.NewNetbox(server, apiKey, deviceType, deviceTypeID)
}
