package jsonrpc2

import (
	"encoding/json"
)

func successResponse(id *json.RawMessage, result interface{}) *Response {
	return &Response{
		JsonRPC: "2.0",
		Result:  result,
		Id:      id,
	}
}

func errorResponse(id *json.RawMessage, err *Error) *Response {
	return &Response{
		JsonRPC: "2.0",
		Error:   err,
		Id:      id,
	}
}
