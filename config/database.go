package config

import (
	"errors"
	"fmt"
)

// Database - Database settings.
type Database struct {
	Name     string `json:"name" yaml:"name"`
	Port     int    `json:"port" yaml:"port"`
	User     string `json:"user" yaml:"user"`
	Password string `json:"password" yaml:"password"`
	Host     string `json:"host" yaml:"host"`
	SslMode  string `json:"sslMode" yaml:"sslMode"`
	URL      string `json:"url" yaml:"url"`
}

// ValidateDatabase - Validates the database settings.
func (d Database) ValidateDatabase() (valid bool) {
	valid = true

	switch "" {
	case d.Host:
	case d.Name:
	case d.User:
	case d.Password:
		valid = false
	}
	if d.Port == 0 {
		valid = false
	}
	return
}

// GetURL parses the database settigns to a string url.
func (d Database) GetURL() (url string, err error) {
	if d.ValidateDatabase() {
		url = fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s port=%d sslmode=%s",
			d.User,
			d.Password,
			d.Name,
			d.Host,
			d.Port,
			d.SslMode,
		)
	} else {
		err = errors.New("failed to parse database url: invalid database data")
	}

	return
}
