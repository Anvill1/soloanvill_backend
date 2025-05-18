package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatus)
	_, err = fmt.Fprintln(w, string(res))
	if err != nil {
		log.Errorln(err)
	}
}

func WrapOK(w http.ResponseWriter, m map[string]interface{}) {
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintln(w, string(res))
	if err != nil {
		log.Errorln(err)
	}
}

func WrapOKWithStatus(w http.ResponseWriter, httpStatus int, responseData string) {
	var m = map[string]string{
		"data":   responseData,
		"result": "OK",
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)
	_, err := fmt.Fprintln(w, string(res))
	if err != nil {
		log.Errorln(err)
	}
}
