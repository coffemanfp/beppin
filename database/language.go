package database

import (
	"fmt"

	"github.com/coffemanfp/beppin-server/database/models"
	dbu "github.com/coffemanfp/beppin-server/database/utils"
	errs "github.com/coffemanfp/beppin-server/errors"
)

func (dS defaultStorage) CreateLanguage(language models.Language) (err error) {
	exists, err := dS.ExistsLanguage(language)
	if err != nil {
		return
	}

	if exists {
		err = fmt.Errorf("failed to check (%s) language: %w", language.Code, errs.ErrExistentObject)
		return
	}

	err = dbu.InsertLanguage(dS.db, language)
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
