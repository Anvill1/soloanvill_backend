package handlers

import (
	"encoding/json"
	"hello/internals/app/models"
	"hello/internals/app/processors"
	"net/http"
)

type UsersHandler struct {
	processor *processors.UserProcessor
}

func NewUserHandler(processor *processors.UserProcessor) *UsersHandler {
	handler := new(UsersHandler)
	handler.processor = processor
	return handler
}

func (handler *UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	var NewJobProcessor processors.JobProcessor

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		WrapError(w, err)
		return
	}

	err = handler.processor.CreateUser(newUser, NewJobProcessor)
	if err != nil {
		WrapError(w, err)
		return
	}

	/*
		var m = map[string]interface{}{
			"result": "OK",
			"data":   "",
		}

	WrapOK(w, m)
	*/
	
	response := newUser.Username + " created"
	WrapOKWithStatus(w, http.StatusCreated, response)
}
