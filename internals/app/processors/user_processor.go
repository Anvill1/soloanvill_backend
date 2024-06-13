package processors

import (
	"errors"
	"hello/internals/app/db"
	"hello/internals/app/models"
	"net/mail"
)

type UserProcessor struct {
	storage *db.UsersStorage
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func NewUserProcessor(storage *db.UsersStorage) *UserProcessor {
	processor := new(UserProcessor)
	processor.storage = storage
	return processor
}

func (processor *UserProcessor) CreateUser(newUser models.User, job JobProcessor) error {
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

	jobprocessor := job.NewJobProcessor()
	err := jobprocessor.CreateJob()
	if err != nil {
		return err
	}

	return processor.storage.CreateUser(newUser)
}
