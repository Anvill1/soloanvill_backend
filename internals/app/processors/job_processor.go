package processors

import (
	"errors"
	"net/http"

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

func (job *JobProcessor) NewJobProcessor() *JobProcessor {
	processor := new(JobProcessor)
	processor.endpoint = "jenkins.soloanvill.ru"
	processor.jobname = "soloanvill_redeploy"
	processor.login = "RTAV3D"
	processor.token = "110486e8e549416302005d62f17ee7099e"
	processor.parameter = "STAGENAME"
	processor.parametervalue = "lovi_bykvi"
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
	if resp.StatusCode != http.StatusCreated {
		log.Errorln(resp)
		return errors.New(resp.Status)
	}
	log.Info(resp)
	return nil
}
