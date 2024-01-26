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
	ErrDone              = config.Err("Error marking todo as done")
	ErrTitleExists       = config.Err("Title already exists")
	ErrNoSuchTask        = config.Err("No such task")
	ErrAlreadyDone       = config.Err("Task already done")
)
