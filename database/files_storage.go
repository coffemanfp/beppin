package database

import "github.com/coffemanfp/beppin/models"

// FileStorage reprensents all implementations for file utils.
type FileStorage interface {
	CreateFile(file models.File) (models.File, error)
	GetFile(fileToFind models.File) (models.File, error)
	ExistsFile(file models.File) (bool, error)
	UpdateFile(fileToUpdate, file models.File) (models.File, error)
	DeleteFile(file models.File) (int, error)
}
