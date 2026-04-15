package handlers

import (
	"encoding/json"
	"net/http"
	"soloanvill_backend/internals/app/models"
	"soloanvill_backend/internals/app/processors"
)

type DeployHandler struct {
	processor *processors.DeployProccessor
}

func NewDeployHandler(processor *processors.DeployProccessor) *DeployHandler {
	handler := new(DeployHandler)
	handler.processor = processor
	return handler
}

func (handler *DeployHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	var NewJobProcessor processors.JobProcessor

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.CreateDeploy(newUser, NewJobProcessor)
	if err != nil {
		WrapError(w, err)
		return
	}

	response := newUser.Username + " created"
	WrapOKWithStatus(w, http.StatusCreated, response)
}
