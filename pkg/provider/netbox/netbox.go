package netbox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lukirs95/monika-gosdk/pkg/types"
)

type Netbox struct {
	server       string
	apiKey       string
	deviceTypeID int
	deviceType   types.DeviceType
	devices      []types.Device
}

func NewNetbox(server string, apiKey string, deviceType types.DeviceType, deviceTypeID int) *Netbox {
	return &Netbox{
		server:       server,
		apiKey:       apiKey,
		deviceType:   deviceType,
		deviceTypeID: deviceTypeID,
		devices:      make([]types.Device, 0),
	}
}

func (netbox *Netbox) GetDeviceType() types.DeviceType {
	return netbox.deviceType
}

func (netbox *Netbox) FetchDevices(ctx context.Context) error {
	path := "api/dcim/devices"
	query := fmt.Sprintf("?device_type_id=%d", netbox.deviceTypeID)
	endpoint := fmt.Sprintf("%s/%s/%s", netbox.server, path, query)

	responses, err := netbox.request(ctx, endpoint)
	if err != nil {
		return err
	}

	allResponses := make([]Device, 0)
	for _, response := range responses {
		var resDevices []Device
		if err := json.Unmarshal(response, &resDevices); err != nil {
			return err
		}
		allResponses = append(allResponses, resDevices...)
	}

	for _, resDevice := range allResponses {
		if resDevice.GetIP() != "" {
			newDevice := types.NewDevice(resDevice.GetId(), netbox.deviceType, resDevice.Name)
			newDevice.SetControlIP(resDevice.GetIP())
			netbox.devices = append(netbox.devices, newDevice)
		}
	}

	return nil
}

func (netbox *Netbox) GetDevices() []types.Device {
	return netbox.devices
}

func (netbox *Netbox) request(ctx context.Context, endpoint string) ([]json.RawMessage, error) {
	next := endpoint
	responses := make([]json.RawMessage, 0)
	for next != "" {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("TOKEN %s", netbox.apiKey))

		var response Response
		if err := json.NewDecoder(req.Body).Decode(&response); err != nil {
			return nil, err
		}
		responses = append(responses, response.Results)
		next = response.Next
	}

	return responses, nil
}
