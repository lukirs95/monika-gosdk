package monika

import (
	"encoding/json"
	"fmt"

	"github.com/lukirs95/monika_gosdk/pkg/types"
)

type errorResponse struct {
	ErrorId string `json:"errorId"`
}

func (m *Monika) ErrorPublish(pubError *types.PubError) (string, error) {
	var eResponse errorResponse
	res, err := m.post("/api/error", pubError)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(res, &eResponse); err != nil {
		return "", nil
	}

	return eResponse.ErrorId, nil
}

func (m *Monika) ErrorDelete(errorId string) error {
	return m.delete(fmt.Sprintf("/api/error/%s", errorId))
}
