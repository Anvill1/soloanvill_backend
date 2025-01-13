package handlers

import (
	"hello/internals/app/processors"
	"net/http"
)

type HealthHandler struct {
	processor *processors.HealthProcessor
}

func NewHealthHandler(processor *processors.HealthProcessor) *HealthHandler {
	handler := new(HealthHandler)
	handler.processor = processor
	return handler
}

func (handler *HealthHandler) Status(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": "up",
	}
	WrapOK(w, response)
}
