package database

import (
	dbu "github.com/coffemanfp/beppin/database/utils"
	"github.com/coffemanfp/beppin/models"
)

func (dS defaultStorage) CreateFile(file models.File) (createdFile models.File, err error) {
	createdFile, err = dbu.InsertFile(dS.db, file)
	return
}

func (dS defaultStorage) GetFile(fileToFind models.File) (file models.File, err error) {
	file, err = dbu.SelectFile(dS.db, fileToFind)
	return
}

// ExistsFile - Exists a file register.
func (dS defaultStorage) ExistsFile(file models.File) (exists bool, err error) {
	exists, err = dbu.ExistsFile(dS.db, file)
	return
}

func (dS defaultStorage) UpdateFile(fileToUpdate, file models.File) (fileUpdated models.File, err error) {
	fileUpdated, err = dbu.UpdateFile(dS.db, fileToUpdate, file)
	return
}

// DeleteFile - Deletes a file register.
func (dS defaultStorage) DeleteFile(file models.File) (id int, err error) {
	id, err = dbu.DeleteFile(dS.db, file)
	return
}
