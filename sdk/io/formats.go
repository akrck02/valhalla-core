package inout

import (
	"encoding/json"
	"io"
)

func ParseJSON[T any](body *io.ReadCloser, object *T) error {
	err := json.NewDecoder(*body).Decode(object)
	if nil != err {
		return err
	}

	return nil
}
