package jsonrpc2

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func writeResponseJson(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (s *Server) handleOneRequest(ctx context.Context, body []byte) *Response {
	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		return errorResponse(nil, ErrParseError("Parse Error"))
	}

	return s.HandleRequest(ctx, req)
}

func (s *Server) handleRequestBatch(ctx context.Context, body []byte) []Response {
	var reqs []Request
	if err := json.Unmarshal(body, &reqs); err != nil {
		return []Response{
			*errorResponse(nil, ErrParseError("Parse error")),
		}
	}

	if len(reqs) == 0 {
		return []Response{
			*errorResponse(nil, ErrInvalidRequest("Invalid Request")),
		}
	}

	responses := make([]Response, 0)

	for _, req := range reqs {
		res := s.HandleRequest(ctx, req)
		if res != nil {
			responses = append(responses, *res)
		}
	}

	return responses
}

func (s *Server) ServeHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		writeResponseJson(w, http.StatusOK, errorResponse(nil, ErrParseError("Parse error")))
	}

	body = bytes.TrimSpace(body)

	if len(body) == 0 {
		writeResponseJson(w, http.StatusOK, errorResponse(nil, ErrInvalidRequest("Invalid request")))
	}

	ctx := r.Context()

	switch body[0] {
	case '{':
		{
			res := s.handleOneRequest(ctx, body)
			if res == nil {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			writeResponseJson(w, http.StatusOK, res)
		}

	case '[':
		{
			responses := s.handleRequestBatch(ctx, body)
			if len(responses) == 0 {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			writeResponseJson(w, http.StatusOK, responses)
		}

	default:
		writeResponseJson(w, http.StatusOK, errorResponse(nil, ErrParseError("Parse error")))
	}
}
