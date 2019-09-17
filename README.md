# go-models-redis
`go-models-redis` its lite and easy model. 

## Requirements
    * Go 1.12 or higher.
    * [go-redis/redis](https://github.com/go-redis/redis) package

## Go-Module
create `go.mod` file in your package folder, and fill below
```
module github.com/eehsiao/go-models-example

go 1.13

require (
	github.com/eehsiao/go-models-lib latest
	github.com/eehsiao/go-models-redis latest
	github.com/go-redis/redis v6.15.5+incompatible
)

```

## Docker
Easy to start the test evn. That you can run the example code.
```bash
$ docker-compose up -d
```

## Usage
```go
import (
    "database/sql"
	"fmt"

	mysql "github.com/eehsiao/go-models-mysql"
	redis "github.com/eehsiao/go-models-redis"
)

// User : json struct that to store into redis
type User struct {
	Host       string `json:"host"`
	User       string `json:"user"`
	SelectPriv string `json:"select_priv"`
}

//new redis dao
redUserModel := &RedUserModel{
    Dao: redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB(),
}

// set a struct for dao as default model (option)
// (*User)(nil) : nil pointer of the User struct
// "user" : is real table name in the db
SetDefaultModel((*User)(nil), "user")
```

## Example
### 1 build-in
[example.go](https://github.com/eehsiao/go-models/blob/master/example/example.go)

The example will connect to local mysql and get user data.
Then connect to local redis and set user data, and get back.

### 2 example
`https://github.com/eehsiao/go-models-example/`


## How-to 
How to design model data logical
### Redis
#### 1.
create a data struct, and add the tag `json:"name"`
```go
type User struct {
	Host       string `json:"host"`
	User       string `json:"user"`
	SelectPriv string `json:"select_priv"`
	IntVal     int    `json:"user,string"`
}
```

if you have integer value, you can add a transfer type desc.
such as json:"user,`string`"

#### 2.
create redis dao
```go
m := redis.NewDao().SetConfig("127.0.0.1:6379", "", 0).OpenDB()
```

#### 3.
directly use the go-redis command function
```go
redBool, err = m.HSet(userTable, redKey, serialStr).Result()
```
