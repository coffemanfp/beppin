package database

import (
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	errs "github.com/coffemanfp/beppin-server/errors"
)

func (dS defaultStorage) CreateLanguage(language models.Language) (id int, err error) {
	exists, err := dS.ExistsLanguage(language)
	if err != nil {
		return
	}

	if exists {
		err = fmt.Errorf("failed to create (%v) language: %w (language)", language.Code, errs.ErrExistentObject)
		return
	}

	id, err = dbu.InsertLanguage(dS.db, language)
	return
}

func (dS defaultStorage) ExistsLanguage(language models.Language) (exists bool, err error) {
	exists, err = dbu.ExistsLanguage(dS.db, language)
	return
}

func (dS defaultStorage) GetLanguage(languageToFind models.Language) (Language models.Language, err error) {
	Language, err = dbu.SelectLanguage(dS.db, languageToFind)
	return
}
