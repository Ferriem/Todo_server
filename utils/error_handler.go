package utils

import (
	"github.com/Ferriem/Todo_server/config"
)

const (
	NoError              = config.Err("")
	ErrUserNotFound      = config.Err("User not found")
	ErrUserExists        = config.Err("User already exists")
	ErrPasswordIncorrect = config.Err("Password incorrect")
	ErrServer            = config.Err("Server error")
	ErrAdd               = config.Err("Error adding todo")
	ErrDelete            = config.Err("Error deleting todo")
	ErrUpdate            = config.Err("Error updating todo")
)
