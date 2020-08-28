package database

import "github.com/coffemanfp/beppin-server/database/models"

// LanguageStorage reprensents all implementations for language utils.
type LanguageStorage interface {
	CreateLanguage(language models.Language) error
	GetLanguage(languageToFind models.Language) (models.Language, error)
	ExistsLanguage(language models.Language) (bool, error)
}