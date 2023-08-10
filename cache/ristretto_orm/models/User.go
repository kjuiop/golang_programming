package models

import (
	"fmt"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Email    string
}

func (u *User) cost() int64 {
	return int64(len(u.Name) + len(u.Password) + len(u.Email) + 8)
}

func (u *User) TableName() string {
	return "users"
}

func (engn *Engn) FindUser(id int64) (*User, bool, error) {
	cacheKey := fmt.Sprintf("user:%d", id)
	if user, has := engn.cache.Get(cacheKey); has {
		if u, ok := user.(User); ok {
			return &u, true, nil
		}
	}

	u := User{Id: id}
	has, err := engn.db.Get(&u)
	if err != nil {
		return nil, false, err
	}

	if !has {
		return nil, false, nil
	}

	engn.cache.SetWithTTL(cacheKey, u, u.cost(), 15*time.Minute)
	return &u, false, nil
}
