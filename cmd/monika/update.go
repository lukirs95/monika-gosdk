package monika

import "github.com/lukirs95/monika-gosdk/pkg/types"

func (m *Monika) UpdatePublish(device *types.Device) error {
	_, err := m.post("/api/update", device)
	return err
}
