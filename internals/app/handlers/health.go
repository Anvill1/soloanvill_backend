package handlers

import (
	"net/http"
	"soloanvill_backend/internals/app/processors"
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
