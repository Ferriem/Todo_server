package utils

import (
	"context"
	"encoding/json"

	"github.com/Ferriem/Todo_server/config"
	"golang.org/x/crypto/bcrypt"
)

type key_value_pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Done  bool   `json:"done"`
}

func NewUser(id string, password string) config.Err {
	client := config.Rdb
	ctx := context.Background()

	value, err := client.Exists(ctx, id).Result()
	if err != nil {
		return ErrServer
	}
	if value == 1 {
		return ErrUserExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ErrServer
	}

	err = client.HSet(ctx, id, id, hashedPassword).Err()
	if err != nil {
		return ErrServer
	}
	return NoError
}

func Login(id string, password string) (bool, config.Err) {
	client := config.Rdb
	ctx := context.Background()

	err := client.HExists(ctx, id, id).Err()
	if err != nil {
		return false, ErrUserNotFound
	}

	value, err := client.HGet(ctx, id, id).Result()
	if err != nil {
		return false, ErrServer
	}

	err = bcrypt.CompareHashAndPassword([]byte(value), []byte(password))
	if err != nil {
		return false, ErrPasswordIncorrect
	}
	return true, NoError
}

func Add(id string, title string, description string) config.Err {
	client := config.Rdb
	ctx := context.Background()
	task := id + "_" + title + "_" + description
	value, err := client.Exists(ctx, task).Result()
	if err != nil {
		return ErrAdd
	}
	if value == 1 {
		return ErrTitleExists
	}
	err = client.HSet(ctx, task, "done", false).Err()
	if err != nil {
		return ErrAdd
	}
	key_value := key_value_pair{
		Key:   title,
		Value: description,
		Done:  false,
	}
	list := id + "_list"
	serializedData, err := json.Marshal(key_value)
	if err != nil {
		return ErrAdd
	}
	err = client.RPush(ctx, list, serializedData).Err()
	if err != nil {
		return ErrAdd
	}
	return NoError
}

func GetFirst(id string) (string, config.Err) {
	client := config.Rdb
	ctx := context.Background()
	list := id + "_list"
	value, err := client.LRange(ctx, list, 0, 0).Result()
	if err != nil {
		return "", ErrServer
	}
	return value[0], NoError
}

func GetInfo(id string) ([]string, config.Err) {
	client := config.Rdb
	ctx := context.Background()
	list := id + "_list"
	value, err := client.LRange(ctx, list, 0, -1).Result()
	if err != nil {
		return nil, ErrServer
	}
	return value, NoError
}

func Done(id string, title string, description string) config.Err {
	client := config.Rdb
	ctx := context.Background()
	task := id + "_" + title + "_" + description
	list := id + "_list"
	value, err := client.Exists(ctx, task).Result()
	if err != nil {
		return ErrDone
	}
	if value == 0 {
		return ErrNoSuchTask
	}
	values, err := client.HGet(ctx, task, "done").Result()
	if err != nil {
		return ErrDone
	}
	if values == "1" {
		return ErrAlreadyDone
	}

	err = client.HSet(ctx, task, "done", true).Err()
	if err != nil {
		return ErrDone
	}
	key_value := key_value_pair{
		Key:   title,
		Value: description,
		Done:  false,
	}
	serializedData, err := json.Marshal(key_value)
	if err != nil {
		return ErrDone
	}
	err = client.LRem(ctx, list, 0, serializedData).Err()
	if err != nil {
		return ErrDone
	}
	key_value.Done = true
	serializedData, err = json.Marshal(key_value)
	if err != nil {
		return ErrDone
	}
	err = client.RPush(ctx, list, serializedData).Err()
	if err != nil {
		return ErrDone
	}
	return NoError
}

func Delete(id string, title string, description string) config.Err {
	client := config.Rdb
	ctx := context.Background()
	list := id + "_list"
	task := id + "_" + title + "_" + description
	value, err := client.Exists(ctx, task).Result()
	if err != nil {
		return ErrDelete
	}
	if value == 0 {
		return ErrNoSuchTask
	}
	done, err := client.HGet(ctx, task, "done").Result()
	if err != nil {
		return ErrDelete
	}
	key_value := key_value_pair{
		Key:   title,
		Value: description,
	}

	if done == "0" {
		key_value.Done = false
	} else {
		key_value.Done = true
	}

	serializedData, err := json.Marshal(key_value)
	if err != nil {
		return ErrDelete
	}

	err = client.LRem(ctx, list, 0, serializedData).Err()
	if err != nil {
		return ErrDelete
	}
	err = client.Del(ctx, task).Err()
	if err != nil {
		return ErrDelete
	}
	return NoError
}
