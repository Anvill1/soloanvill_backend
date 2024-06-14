package processors

import (
	"errors"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type JobProcessor struct {
	endpoint       string
	jobname        string
	login          string
	token          string
	parameter      string // Наименование параметра для передачи в Jenkins job'у
	parametervalue string // Значение параметра для передачи в Jenkins job'у
	url            string
}

func (job *JobProcessor) NewJobProcessor(endpoint string, login string, token string, jobname string, parameter string, parametervalue string) *JobProcessor {
	processor := new(JobProcessor)
	processor.endpoint = endpoint
	processor.login = login
	processor.token = token
	processor.jobname = jobname
	processor.parameter = parameter
	processor.parametervalue = url.QueryEscape(parametervalue)
	processor.url = "https://" + processor.endpoint + "/job/" + processor.jobname + "/buildWithParameters?" + processor.parameter + "=" + processor.parametervalue
	return processor
}

func (processor *JobProcessor) CreateJob() error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", processor.url, nil)
	if err != nil {
		log.Errorln(err)
		return err
	}

	req.SetBasicAuth(processor.login, processor.token)

	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		log.Errorln(resp)
		return errors.New(resp.Status)
	}
	log.Info(resp)
	return nil
}
