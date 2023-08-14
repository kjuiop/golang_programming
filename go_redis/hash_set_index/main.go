package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type User struct {
	Id       int64  `redis:"id"`
	Name     string `redis:"name"`
	Email    string `redis:"email"`
	Password string `redis:"password"`
	Age      int64  `redis:"age"`
}

func main() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6384",
		Password: "",
		DB:       0,
	})
	users := []User{
		{Id: 1, Name: "Jake1", Email: "arneg0shua@gmail.com", Password: "111", Age: 21},
		{Id: 2, Name: "Jake2", Email: "arneg1shua@gmail.com", Password: "222", Age: 22},
		{Id: 3, Name: "Jake3", Email: "arneg2shua@gmail.com", Password: "333", Age: 23},
		{Id: 4, Name: "Jake4", Email: "arneg3shua@gmail.com", Password: "444", Age: 24},
		{Id: 5, Name: "Jake5", Email: "arneg4shua@gmail.com", Password: "555", Age: 25},
	}
	for i := 0; i < len(users); i++ {
		if err := CreateUser(rdb, ctx, users[i]); err != nil {
			panic(err)
		}
	}
	time.Sleep(100 * time.Millisecond)
	u1, err := FindUserByEmail(rdb, ctx, "arneg0shua@gmail.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("1st: %v\n", u1)

	u2, err := FindUserByAge(rdb, ctx, 23, 25)
	if err != nil {
		panic(err)
	}
	spew.Dump("2nd: %v\n", u2)

}

func CreateUser(rdb *redis.Client, ctx context.Context, u User) error {
	cmds, err := rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		// data 등록
		if err := pipe.HSet(ctx, fmt.Sprintf("users:%d", u.Id), u).Err(); err != nil {
			return err
		}
		// index 등록
		if err := pipe.Set(ctx, fmt.Sprintf("users:email:%s", u.Email), u.Id, 0).Err(); err != nil {
			return err
		}
		// index 등록
		if err := pipe.ZAdd(ctx, "users:age", redis.Z{Score: float64(u.Age), Member: u.Id}).Err(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	for i := 0; i < len(cmds); i++ {
		fmt.Printf("%d-%d: %v\n", u.Id, i, cmds[i].Args())
	}
	return nil
}

func FindUser(rdb *redis.Client, ctx context.Context, id int64) (*User, error) {
	res, err := rdb.HGetAll(ctx, fmt.Sprintf("users:%d", id)).Result()
	if err != nil {
		return nil, err
	}

	user := new(User)
	user.Id, err = strconv.ParseInt(res["id"], 10, 64)
	if err != nil {
		return nil, err
	}
	user.Age, err = strconv.ParseInt(res["age"], 10, 64)
	if err != nil {
		return nil, err
	}
	user.Name = res["name"]
	user.Email = res["email"]
	user.Password = res["password"]

	return user, nil
}

func FindUserByEmail(rdb *redis.Client, ctx context.Context, email string) (*User, error) {
	res, err := rdb.Get(ctx, "users:email:"+email).Int64()
	if err != nil {
		return nil, err
	}
	user, err := FindUser(rdb, ctx, res)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindUserByAge(rdb *redis.Client, ctx context.Context, startAge, endAge int64) ([]*User, error) {
	res, err := rdb.ZRangeByScore(ctx, "users:age", &redis.ZRangeBy{Min: strconv.FormatInt(startAge, 10), Max: strconv.FormatInt(endAge, 10)}).Result()
	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)
	for i := 0; i < len(res); i++ {
		id, err := strconv.ParseInt(res[i], 10, 64)
		if err != nil {
			return nil, err
		}
		user, err := FindUser(rdb, ctx, id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
