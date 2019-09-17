// Author :		Eric<eehsiao@gmail.com>

package main

import (
	"fmt"

	lib "github.com/eehsiao/go-models-lib"
	redis "github.com/eehsiao/go-models-redis"
)

var (
	redDao    *redis.Dao
	user      *User
	serialStr string
	redKey    string
	redBool   bool
	err       error
)

func main() {

	redUserModel := &RedUserModel{
		Dao: redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB(),
	}

	if err = redUserModel.SetDefaultModel((*User)(nil), "user"); err != nil {
		panic(err.Error())
	}

	user = &User{
		Host:       "test1",
		User:       "test2",
		SelectPriv: "Y",
	}

	if serialStr, err = lib.Serialize(user); err == nil {
		redKey = user.Host + user.User
		// HSet is github.com/go-redis/redis original command
		if redBool, err = redUserModel.HSet("user", redKey, serialStr).Result(); err != nil {
			panic(err.Error())
		}
	}

	if user, err = redUserModel.UserHGet(redKey); err == nil {
		fmt.Println(fmt.Sprintf("UserHGet : %s = %v", redKey, user))
	}
}
