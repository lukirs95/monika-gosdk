package provider

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/lukirs95/monika-gosdk/pkg/types"
)

func TestNetbox(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		t.Fatal(err)
	}
	endpoint := os.Getenv("NETBOX_ADDRESS")
	apikey := os.Getenv("NETBOX_APIKEY")
	deviceType := types.DeviceType(os.Getenv("NETBOX_DEVICETYPE"))
	deviceTypeId, err := strconv.Atoi(os.Getenv("NETBOX_DEVICETYPE_ID"))
	if err != nil {
		t.Fatal(err)
	}
	provider := NewDeviceProviderNetbox(endpoint, apikey, deviceType, deviceTypeId)
	if err := provider.FetchDevices(context.Background()); err != nil {
		t.Fatal(err)
	}

	t.Log(provider.GetDevices())
}
