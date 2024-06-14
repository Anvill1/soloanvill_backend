package processors

import (
	"errors"
	"hello/internals/app/db"
	"hello/internals/app/models"
	"hello/internals/cfg"
	"net/mail"
)

type DeployProccessor struct {
	storage *db.DeployStorage
	cfg     *cfg.Cfg
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NewDeployProccessor(storage *db.DeployStorage, cfg *cfg.Cfg) *DeployProccessor {
	processor := new(DeployProccessor)
	processor.storage = storage
	processor.cfg = cfg
	return processor
}

func (processor *DeployProccessor) CreateDeploy(newUser models.User, job JobProcessor) error {
	if newUser.Username == "" {
		//log.Errorln("username should not be empy")
		return errors.New("username should not be empty")
	}
	if newUser.Email == "" {
		//log.Errorln("email should not be empty")
		return errors.New("email should not be empty")
	}

	if !valid(newUser.Email) {
		//log.Errorln("Fail valid email")
		return errors.New("email is incorrect")
	}
	jobprocessor := job.NewJobProcessor(processor.cfg.JenkinsEndpoint, processor.cfg.JenkinsLogin, processor.cfg.JenkinsToken, "soloanvill_redeploy", "USERNAME", newUser.Username)
	err := jobprocessor.CreateJob()
	if err != nil {
		return err
	}
	return processor.storage.CreateDeploy(newUser)
}
