package utils

import (
	"context"

	"github.com/Ferriem/Todo_server/config"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type key_value_pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
	key_value := key_value_pair{
		Key:   title,
		Value: description,
	}
	list := id + "_list"
	err := client.RPush(ctx, list, key_value).Err()
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

func Delete(id string) config.Err {
	client := config.Rdb
	ctx := context.Background()
	list := id + "_list"
	err := client.Del(ctx, list).Err()
	if err != nil {
		return ErrDelete
	}
	return NoError
}

func Update(id string, title string, description string) config.Err {
	client := config.Rdb
	ctx := context.Background()
	key_value := key_value_pair{
		Key:   title,
		Value: description,
	}
	list := id + "_list"
	err := client.LSet(ctx, list, 0, key_value).Err()
	if err != nil {
		return ErrUpdate
	}
	return NoError
}
