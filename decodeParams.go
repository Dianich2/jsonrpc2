package jsonrpc2

import (
	"encoding/json"
	"errors"
)

func DecodeParams[T any](raw json.RawMessage) (T, error) {
	var curVar T

	// if len(raw) == 0 {
	// 	return curVar, errors.New("params are required")
	// }

	if err := json.Unmarshal(raw, &curVar); err != nil {
		return curVar, err
	}

	return curVar, nil
}

func DecodeParamsInto(raw json.RawMessage, dest interface{}) error {
	// if len(raw) == 0 {
	// 	return errors.New("params are required")
	// }

	if dest == nil {
		return errors.New("dest cannot be nil")
	}

	return json.Unmarshal(raw, dest)
}
