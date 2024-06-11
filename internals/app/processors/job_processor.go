package processors

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type JobProcessor struct {
	endpoint       string
	name           string
	token          string
	parameter      string //? добавить jobparameter string под наименование параметра, который будет ловить Jenkins
	parametervalue string //? Добавить username string под параметры имени, которые нужно будет ловить из UI Flask'а
	url            string
}

func (job *JobProcessor) NewJobProcessor() *JobProcessor {
	processor := new(JobProcessor)
	job.endpoint = "jenkins.soloanvill.ru"
	job.name = "soloanvill_redeploy"
	job.token = "34fK9SC11im2UY6h"
	job.parameter = "STAGENAME"
	job.parametervalue = "Prod"
	processor.url = "https://" + job.endpoint + "/job/" + job.name + "/buildWithParameters?token=" + job.token + "&" + job.parameter + "=" + job.parametervalue
	return processor
}

func (job *JobProcessor) CreateJob() error {

	resp, err := http.Post(job.url, "application/json", nil)
	if err != nil {
		log.Errorln(err)
		return err
	}
	if resp.StatusCode != 200 {
		log.Errorln(resp)
		return errors.New(resp.Status)
	}
	log.Info(resp)
	return nil
}
