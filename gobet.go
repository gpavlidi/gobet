package gobet

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
)

type Session struct {
	token         string
	configuration *Configuration
	logger        *log.Logger
}

type Configuration struct {
	Username string
	Password string
	AppKey   string
	Locale   string
	LogDest  io.Writer
}

func NewSession(c *Configuration) (*Session, error) {
	if c.Username == "" {
		return nil, errors.New("Username cannot be empty!")
	}
	if c.Password == "" {
		return nil, errors.New("Password cannot be empty!")
	}
	if c.Locale == "" {
		c.Locale = "en"
	}
	if c.LogDest == nil {
		c.LogDest = ioutil.Discard
	}

	logger := log.New(c.LogDest, "", log.LstdFlags)
	s := Session{configuration: c, logger: logger}

	return &s, nil
}
