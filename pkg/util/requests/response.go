package requests

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Response struct {
	Text []byte
}

func (res *Response) JSON() (map[string]string, error) {
	// TODO: maybe use interface instead of string?
	var data map[string]string
	if err := json.Unmarshal(res.Text, &data); err != nil {
		return data, errors.Wrap(err, "error unmarshal json using map[string]string")
	}
	return data, nil
}
