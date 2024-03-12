package monika

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

type Monika struct {
	client   *http.Client
	endpoint string
}

func NewMonika(endpoint string) *Monika {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	return &Monika{
		client:   client,
		endpoint: endpoint,
	}
}

func (m *Monika) post(path string, body any) ([]byte, error) {
	endpoint := fmt.Sprintf("%s%s", m.endpoint, path)
	reqBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(reqBody)

	res, err := m.client.Post(endpoint, "application/json", bodyReader)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status: %d, %s", res.StatusCode, res.Status)
	}

	return io.ReadAll(res.Body)
}

func (m *Monika) delete(path string) error {
	endpoint := fmt.Sprintf("%s%s", m.endpoint, path)
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return err
	}

	res, err := m.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("status: %d, %s", res.StatusCode, res.Status)
	}

	return nil
}
