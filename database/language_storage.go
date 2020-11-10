package database

import "github.com/coffemanfp/beppin/models"

// LanguageStorage reprensents all implementations for language utils.
type LanguageStorage interface {
	CreateLanguage(language models.Language) (int, error)
	GetLanguage(languageToFind models.Language) (models.Language, error)
	ExistsLanguage(language models.Language) (bool, error)
}
