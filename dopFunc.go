package jsonrpc2

import (
	"context"
	"encoding/json"
	"runtime/debug"
)

func validateRequest(req Request) *Error {
	if req.JsonRPC != "2.0" {
		return ErrInvalidRequest(`jsonrpc must be "2.0"`)
	}

	if req.Method == "" {
		return ErrInvalidRequest("method must not be empty")
	}

	return nil
}

func (s *Server) callHandlerSafely(ctx context.Context, handler HandlerFunc, params json.RawMessage) (res interface{}, err *Error) {
	defer func() {
		if r := recover(); r != nil {
			if s.logger != nil {
				s.logger.Printf("jsonrpc: panic recovered: %v\n%s", r, debug.Stack())
			}
			res = nil
			err = ErrInternalError("Internal error")
		}
	}()

	return handler(ctx, params)
}
