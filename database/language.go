package database

import (
	"fmt"

	dbu "github.com/coffemanfp/beppin/database/utils"
	errs "github.com/coffemanfp/beppin/errors"
	"github.com/coffemanfp/beppin/models"
)

func (dS defaultStorage) CreateLanguage(language models.Language) (createdLanguage models.Language, err error) {
	exists, err := dS.ExistsLanguage(language)
	if err != nil {
		return
	}

	if exists {
		err = fmt.Errorf("failed to create (%v) language: %w (language)", language.Code, errs.ErrExistentObject)
		return
	}

	createdLanguage, err = dbu.InsertLanguage(dS.db, language)
	return
}

func (dS defaultStorage) ExistsLanguage(language models.Language) (exists bool, err error) {
	exists, err = dbu.ExistsLanguage(dS.db, language)
	return
}

func (dS defaultStorage) GetLanguage(languageToFind models.Language) (language models.Language, err error) {
	language, err = dbu.SelectLanguage(dS.db, languageToFind)
	return
}
