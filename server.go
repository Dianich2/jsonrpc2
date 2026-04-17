package jsonrpc2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
)

type HandlerFunc func(ctx context.Context, params json.RawMessage) (interface{}, *Error)

type Server struct {
	mutex   sync.RWMutex
	methods map[string]HandlerFunc
	logger  Logger
}

// конструктор для удобного создания экземпляра сервера
func New() *Server {
	return &Server{
		methods: make(map[string]HandlerFunc),
	}
}

// функция для регистрации метода
func (s *Server) Register(methodName string, handler HandlerFunc) error {
	if methodName == "" {
		return errors.New("jsonrpc: method name cannot be empty")
	}

	if handler == nil {
		return errors.New("jsonrpc: handler cannot be nil")
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, exists := s.methods[methodName]; exists {
		return fmt.Errorf("jsonrpc: method %q is already exists", methodName)
	}

	s.methods[methodName] = handler
	return nil
}

// приватная функция для поиска метода сервером
func (s *Server) searchMethodByName(methodName string) (HandlerFunc, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	handler, res := s.methods[methodName]
	return handler, res
}

// функция обработки запроса
func (s *Server) HandleRequest(ctx context.Context, req Request) *Response {
	isNotification := req.Id == nil

	if err := validateRequest(req); err != nil {
		if isNotification {
			return nil
		}

		return errorResponse(req.Id, err)
	}

	curHandler, isFound := s.searchMethodByName(req.Method)

	if !isFound {
		if isNotification {
			return nil
		}

		return errorResponse(req.Id, ErrMethodNotFound(req.Method))
	}

	res, err := s.callHandlerSafely(ctx, curHandler, req.Params)
	if err != nil {
		if isNotification {
			return nil
		}

		return errorResponse(req.Id, err)
	}

	if isNotification {
		return nil
	}

	return successResponse(req.Id, res)
}
