package netbox

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lukirs95/monika_gosdk/pkg/types"
)

type Netbox struct {
	server       string
	apiKey       string
	deviceTypeID int
}

func NewNetbox(server string, apiKey string, deviceTypeID int) *Netbox {
	return &Netbox{
		server:       server,
		apiKey:       apiKey,
		deviceTypeID: deviceTypeID,
	}
}

func (netbox *Netbox) GetDevices(ctx context.Context) ([]types.Device, error) {
	path := "api/dcim/devices"
	query := fmt.Sprintf("?device_type_id=%d", netbox.deviceTypeID)
	endpoint := fmt.Sprintf("%s/%s/%s", netbox.server, path, query)

	responses, err := netbox.request(ctx, endpoint)
	if err != nil {
		return nil, err
	}

	allResponses := make([]Device, 0)
	for _, response := range responses {
		var resDevices []Device
		if err := json.Unmarshal(response, &resDevices); err != nil {
			return nil, err
		}
		allResponses = append(allResponses, resDevices...)
	}

	parsedDevices := make([]types.Device, 0)
	for _, resDevice := range allResponses {
		if resDevice.GetIP() != "" {
			parsedDevices = append(parsedDevices, types.Device{
				Name:      resDevice.GetName(),
				ControlIP: resDevice.GetIP(),
			})
		}
	}

	return parsedDevices, nil
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
