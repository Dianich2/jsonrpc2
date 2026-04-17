//omitempty означает, что оно будет включено в json только, если оно не пустое

package jsonrpc2

import "encoding/json"

type Request struct {
	JsonRPC string           `json:"jsonrpc"`
	Method  string           `json:"method"`
	Params  json.RawMessage  `json:"params,omitempty"`
	Id      *json.RawMessage `json:"id,omitempty"`
}

type Response struct {
	JsonRPC string           `json:"jsonrpc"`
	Result  interface{}      `json:"result,omitempty"`
	Error   *Error           `json:"error,omitempty"`
	Id      *json.RawMessage `json:"id"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
