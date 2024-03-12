package monika

import "github.com/lukirs95/monika_gosdk/pkg/types"

func (m *Monika) Authenticate(authObject *types.Auth) error {
	_, err := m.post("/api/authenticate", authObject)
	if err != nil {
		return err
	}

	return nil
}
